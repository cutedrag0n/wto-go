package transport

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func testClient(t *testing.T, h http.HandlerFunc, opts func(*Config)) *Client {
	t.Helper()
	s := httptest.NewServer(h)
	t.Cleanup(s.Close)
	c := Config{APIKey: "secret", BaseURL: s.URL, Timeout: 3 * time.Second, Headers: http.Header{}}
	if opts != nil {
		opts(&c)
	}
	cl, err := New(c)
	if err != nil {
		t.Fatal(err)
	}
	return cl
}
func TestErrors(t *testing.T) {
	for _, status := range []int{400, 401, 403, 404, 422, 429, 500, 502, 503, 504} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				w.Write([]byte(`{"title":"secret","errors":{"i":["required"]}}`))
			}, nil)
			_, err := c.Do(context.Background(), Request{Service: "qr", Path: "/members", Method: "GET"})
			var api *APIError
			if !errors.As(err, &api) || api.StatusCode != status || api.Errors["i"][0] != "required" {
				t.Fatal(err)
			}
			if strings.Contains(err.Error(), "secret") {
				t.Fatal("credential reflected")
			}
		})
	}
}
func TestRetryAndHeaders(t *testing.T) {
	calls := 0
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Header.Get("Ocp-Apim-Subscription-Key") != "secret" {
			t.Error("authentication")
		}
		if calls == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(503)
			return
		}
		w.Write([]byte(`[]`))
	}, func(c *Config) { c.Retries = 1 })
	if _, err := c.Do(context.Background(), Request{Service: "qr", Path: "/members", Method: "GET"}); err != nil {
		t.Fatal(err)
	}
	if calls != 2 || len(c.config.Limiter.sent) != 2 {
		t.Fatal(calls)
	}
}
func TestDecodeAndNonJSONError(t *testing.T) {
	for _, status := range []int{200, 500} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			c := testClient(t, func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(status); w.Write([]byte("not JSON")) }, nil)
			_, err := Decode[[]int](context.Background(), c, Request{Service: "qr", Path: "/members", Method: "GET"})
			if err == nil {
				t.Fatal("expected error")
			}
			var api *APIError
			if status == 500 && (!errors.As(err, &api) || string(api.Body) != "not JSON") {
				t.Fatal(err)
			}
		})
	}
}
func TestTimeout(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) { <-r.Context().Done() }, func(c *Config) { c.Timeout = 20 * time.Millisecond })
	_, err := c.Do(context.Background(), Request{Service: "qr", Path: "/members", Method: "GET"})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
}
func TestRetryAfter(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)
	for _, tt := range []struct {
		value string
		want  time.Duration
	}{{"3", 3 * time.Second}, {now.Add(5 * time.Second).Format(http.TimeFormat), 5 * time.Second}, {"-1", 0}, {"invalid", 0}, {now.Add(-time.Second).Format(http.TimeFormat), 0}} {
		if got := retryAfter(tt.value, now); got != tt.want {
			t.Fatalf("%s: %v", tt.value, got)
		}
	}
}
func TestQueryEncoding(t *testing.T) {
	q := url.Values{}
	zero := 0
	no := false
	Add(q, "zero", &zero)
	Add(q, "false", &no)
	Add(q, "countries[]", []string{"004", "008"})
	Add(q, "nil", (*int)(nil))
	if q.Get("zero") != "0" || q.Get("false") != "false" || len(q["countries[]"]) != 2 || q.Has("nil") {
		t.Fatal(q)
	}
}

func Test429CooldownAndRetryCancellation(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) { w.Header().Set("Retry-After", "2"); w.WriteHeader(429) }, nil)
	in := Request{Service: "qr", Path: "/members", Method: "GET"}
	if _, err := c.Do(context.Background(), in); err == nil {
		t.Fatal("missing 429")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	if _, err := c.Do(ctx, in); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
	if len(c.config.Limiter.sent) != 1 {
		t.Fatal("cooldown consumed quota")
	}
}

type temporaryError struct{}

func (temporaryError) Error() string   { return "secret" }
func (temporaryError) Timeout() bool   { return false }
func (temporaryError) Temporary() bool { return true }

type fakeRoundTripper func(*http.Request) (*http.Response, error)

func (f fakeRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func TestTransportErrorAndRetry(t *testing.T) {
	calls := 0
	cfg := Config{APIKey: "secret", BaseURL: "https://example.test", Timeout: 3 * time.Second, Retries: 1,
		HTTPClient: &http.Client{Transport: fakeRoundTripper(func(r *http.Request) (*http.Response, error) { calls++; return nil, temporaryError{} })}}
	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Do(context.Background(), Request{Service: "qr", Path: "/members", Method: "GET"})
	var ne temporaryError
	if err == nil || !errors.As(err, &ne) || strings.Contains(err.Error(), "secret") || calls != 2 {
		t.Fatal(err, calls)
	}
	cfg.Timeout = 5 * time.Millisecond
	c, err = New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Do(context.Background(), Request{Service: "qr", Path: "/members", Method: "GET"})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
}

func TestSameOriginRedirectCountsAgainstQuota(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/final" {
			http.Redirect(w, r, "/final", 302)
			return
		}
		w.Write([]byte(`[]`))
	}, nil)
	if _, err := c.Do(context.Background(), Request{Service: "qr", Path: "/members", Method: "GET"}); err != nil {
		t.Fatal(err)
	}
	if len(c.config.Limiter.sent) != 2 {
		t.Fatal("redirect was not counted")
	}
	if d := c.config.Limiter.sent[1].Sub(c.config.Limiter.sent[0]); d < 100*time.Millisecond {
		t.Fatal("redirect exceeded rate", d)
	}
}
