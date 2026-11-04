package wto

import (
	"net/http"
	"time"

	"github.com/cutedrag0n/wto-go/internal/transport"
)

// Option configures a new Client.
type Option func(*transport.Config)

// WithBaseURL overrides the gateway URL, retaining each service's official prefix.
func WithBaseURL(u string) Option { return func(c *transport.Config) { c.BaseURL = u } }

// Service identifies a WTO API for a service-specific URL override.
type Service string

const (
	Timeseries Service = "timeseries"
	EPing      Service = "eping"
	QR         Service = "qr"
	TFAD       Service = "tfad"
)

// WithServiceBaseURL sets the complete base URL including the service prefix.
func WithServiceBaseURL(s Service, u string) Option {
	return func(c *transport.Config) { c.ServiceURLs[string(s)] = u }
}

// WithTimeout sets the total call budget, including quota waiting and retries.
func WithTimeout(d time.Duration) Option { return func(c *transport.Config) { c.Timeout = d } }

// WithRetry sets the number of additional attempts (zero by default).
func WithRetry(n int) Option { return func(c *transport.Config) { c.Retries = n } }

// WithProxy configures an HTTP, HTTPS or SOCKS5 proxy.
func WithProxy(u string) Option { return func(c *transport.Config) { c.Proxy = u } }

// WithHTTPClient copies the client configuration. Transport and Jar remain shared.
func WithHTTPClient(h *http.Client) Option { return func(c *transport.Config) { c.HTTPClient = h } }

// WithHeader adds a default request header. Authentication always uses the constructor's key.
func WithHeader(k, v string) Option { return func(c *transport.Config) { c.Headers.Set(k, v) } }

// WithRateLimiter shares quota state with other clients using the same subscription.
func WithRateLimiter(l *RateLimiter) Option { return func(c *transport.Config) { c.Limiter = l } }
