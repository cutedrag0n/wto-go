// Package timeseries provides the WTO timeseries API. Construct its Service through wto.NewClient.
package timeseries

import (
	"github.com/cutedrag0n/wto-go/internal/transport"
)

// Service uses a shared authenticated transport.
type Service struct{ client *transport.Client }

// NewService is used by the root client to wire this service.
func NewService(c *transport.Client) *Service { return &Service{client: c} }
