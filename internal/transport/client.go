package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

// Config is frozen when the root client is constructed.
type Config struct {
	APIKey      string
	BaseURL     string
	ServiceURLs map[string]string
	Timeout     time.Duration
	Retries     int
	Proxy       string
	HTTPClient  *http.Client
	Headers     http.Header
	Limiter     *RateLimiter
}

type Client struct {
	http   *resty.Client
	config Config
}

func New(c Config) (*Client, error) {
	if strings.TrimSpace(c.APIKey) == "" {
		return nil, errors.New("wto: API key is required")
	}
	if c.Timeout <= 0 || c.Retries < 0 {
		return nil, errors.New("wto: invalid timeout or retry count")
	}
	for service := range c.ServiceURLs {
		switch service {
		case "timeseries", "eping", "qr", "tfad":
		default:
			return nil, errors.New("wto: unknown service")
		}
	}
	for _, u := range append([]string{c.BaseURL}, mapValues(c.ServiceURLs)...) {
		p, err := url.Parse(u)
		if err != nil || (p.Scheme != "http" && p.Scheme != "https") || p.Host == "" || p.User != nil || p.RawQuery != "" || p.Fragment != "" {
			return nil, errors.New("wto: invalid base URL")
		}
	}
	hc := http.Client{}
	if c.HTTPClient != nil {
		hc = *c.HTTPClient
	}
	if c.Proxy != "" {
		p, err := url.Parse(c.Proxy)
		if err != nil || p.Host == "" || (p.Scheme != "http" && p.Scheme != "https" && p.Scheme != "socks5") {
			return nil, errors.New("wto: invalid proxy URL")
		}
		tr := hc.Transport
		if tr == nil {
			tr = http.DefaultTransport
		}
		standard, ok := tr.(*http.Transport)
		if !ok {
			return nil, errors.New("wto: proxy requires *http.Transport")
		}
		clone := standard.Clone()
		clone.Proxy = http.ProxyURL(p)
		hc.Transport = clone
	}
	if c.Limiter == nil {
		c.Limiter = NewRateLimiter()
	}
	c.Headers = c.Headers.Clone()
	urls := make(map[string]string, len(c.ServiceURLs))
	for k, v := range c.ServiceURLs {
		urls[k] = v
	}
	c.ServiceURLs = urls
	originalRedirect := hc.CheckRedirect
	hc.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > 0 && (req.URL.Scheme != via[0].URL.Scheme || req.URL.Host != via[0].URL.Host) {
			return http.ErrUseLastResponse
		}
		if originalRedirect != nil {
			return originalRedirect(req, via)
		}
		if len(via) >= 10 {
			return errors.New("wto: too many redirects")
		}
		return nil
	}
	underlying := hc.Transport
	if underlying == nil {
		underlying = http.DefaultTransport
	}
	hc.Transport = &quotaTransport{base: underlying, limiter: c.Limiter}
	r := resty.NewWithClient(&hc)
	return &Client{http: r, config: c}, nil
}

func mapValues(m map[string]string) []string {
	v := make([]string, 0, len(m))
	for _, s := range m {
		v = append(v, s)
	}
	return v
}

// Request describes one documented operation; services own its query and body semantics.
type Request struct {
	Service, Path, Method string
	Query                 url.Values
	Body                  any
	Data                  bool
}

func (c *Client) Do(ctx context.Context, in Request) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	ctx = context.WithValue(ctx, dataRequestKey{}, in.Data)
	defer cancel()
	base := strings.TrimRight(c.config.BaseURL, "/") + map[string]string{
		"timeseries": "/timeseries/v1", "eping": "/eping", "qr": "/qrs", "tfad": "/tfad",
	}[in.Service]
	if override := c.config.ServiceURLs[in.Service]; override != "" {
		base = strings.TrimRight(override, "/")
	}
	var body []byte
	if in.Body != nil {
		var err error
		body, err = json.Marshal(in.Body)
		if err != nil {
			return nil, fmt.Errorf("wto: encode request: %w", err)
		}
	}
	for attempt := 0; ; attempt++ {
		r := c.http.R().SetContext(ctx).SetQueryParamsFromValues(in.Query)
		for k, vs := range c.config.Headers {
			for _, v := range vs {
				r.Header.Add(k, v)
			}
		}
		r.SetHeader("Ocp-Apim-Subscription-Key", c.config.APIKey)
		if body != nil {
			r.SetHeader("Content-Type", "application/json").SetBody(body)
		}
		res, err := r.Execute(in.Method, base+in.Path)
		if err != nil {
			if ctx.Err() != nil {
				return nil, fmt.Errorf("wto: request canceled: %w", ctx.Err())
			}
			var ne net.Error
			if attempt < c.config.Retries && errors.As(err, &ne) && (ne.Timeout() || ne.Temporary()) {
				if err := sleep(ctx, backoff(attempt)); err != nil {
					return nil, fmt.Errorf("wto: retry: %w", err)
				}
				continue
			}
			// Resty/net errors can contain URLs and headers; preserve matching without printing secrets.
			return nil, &requestError{cause: err}
		}
		status := res.StatusCode()
		if status >= 200 && status < 300 {
			return res.Body(), nil
		}
		delay := retryAfter(res.Header().Get("Retry-After"), time.Now())
		if status == 429 && delay > 0 {
			c.config.Limiter.Cooldown(time.Now().Add(delay))
		}
		retryable := status == 429 || status == 502 || status == 503 || status == 504
		if retryable && attempt < c.config.Retries {
			if err := sleep(ctx, max(delay, backoff(attempt))); err != nil {
				return nil, fmt.Errorf("wto: retry: %w", err)
			}
			continue
		}
		apiErr := &APIError{}
		// Preserve malformed and non-JSON responses even if structured decoding fails.
		if err := json.Unmarshal(res.Body(), apiErr); err != nil {
			apiErr = &APIError{}
		}
		apiErr.StatusCode = status
		apiErr.Header = res.Header().Clone()
		apiErr.Body = append([]byte(nil), res.Body()...)
		return nil, apiErr
	}
}

type requestError struct{ cause error }

func (e *requestError) Error() string { return "wto: HTTP request failed" }
func (e *requestError) Unwrap() error { return e.cause }

func backoff(attempt int) time.Duration {
	return min(time.Second*time.Duration(1<<min(attempt, 5)), 30*time.Second)
}
func retryAfter(value string, now time.Time) time.Duration {
	if seconds, err := strconv.ParseInt(value, 10, 32); err == nil && seconds >= 0 {
		return time.Duration(seconds) * time.Second
	}
	if date, err := http.ParseTime(value); err == nil {
		return max(time.Duration(0), date.Sub(now))
	}
	return 0
}

func Decode[T any](ctx context.Context, c *Client, in Request) (T, error) {
	var result T
	b, err := c.Do(ctx, in)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(b, &result); err != nil {
		return result, fmt.Errorf("wto: decode %s%s: %w", in.Service, in.Path, err)
	}
	return result, nil
}

// Add adds optional query values without reflection. Pointers preserve explicit false and zero.
func Add(q url.Values, key string, value any) {
	switch v := value.(type) {
	case string:
		if v != "" {
			q.Set(key, v)
		}
	case *int:
		if v != nil {
			q.Set(key, strconv.Itoa(*v))
		}
	case *bool:
		if v != nil {
			q.Set(key, strconv.FormatBool(*v))
		}
	case []string:
		for _, s := range v {
			q.Add(key, s)
		}
	}
}

// quotaTransport gates every HTTP exchange, including same-origin redirects.
// Keeping admission at the RoundTripper boundary prevents retry paths bypassing quotas.
type dataRequestKey struct{}
type quotaTransport struct {
	base    http.RoundTripper
	limiter *RateLimiter
}

func (t *quotaTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	data, _ := r.Context().Value(dataRequestKey{}).(bool)
	if err := t.limiter.Wait(r.Context(), data); err != nil {
		return nil, err
	}
	return t.base.RoundTrip(r)
}
