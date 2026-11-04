package wto

import (
	"net/http"
	"time"

	"github.com/cutedrag0n/wto-go/eping"
	"github.com/cutedrag0n/wto-go/internal/transport"
	"github.com/cutedrag0n/wto-go/qr"
	"github.com/cutedrag0n/wto-go/tfad"
	"github.com/cutedrag0n/wto-go/timeseries"
)

// Client groups the WTO services using one transport and subscription limiter.
type Client struct {
	Timeseries *timeseries.Service
	EPing      *eping.Service
	QR         *qr.Service
	TFAD       *tfad.Service
}

// NewClient constructs a concurrent-safe SDK client. Configuration is immutable after construction.
func NewClient(apiKey string, opts ...Option) (*Client, error) {
	c := transport.Config{APIKey: apiKey, BaseURL: "https://api.wto.org", Timeout: 30 * time.Second,
		Headers: http.Header{}, ServiceURLs: map[string]string{}}
	for _, opt := range opts {
		if opt != nil {
			opt(&c)
		}
	}
	t, err := transport.New(c)
	if err != nil {
		return nil, err
	}
	return &Client{Timeseries: timeseries.NewService(t), EPing: eping.NewService(t), QR: qr.NewService(t), TFAD: tfad.NewService(t)}, nil
}
