// Package tfad provides the WTO tfad API. Construct its Service through wto.NewClient.
package tfad

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/cutedrag0n/wto-go/internal/transport"
)

// Service uses a shared authenticated transport.
type Service struct{ client *transport.Client }

// NewService is used by the root client to wire this service.
func NewService(c *transport.Client) *Service { return &Service{client: c} }

// TransparencyParams configures Transparency. Omitted fields retain server defaults.
type TransparencyParams struct {
	// Countries: Country codes are to be sent in the ISO 3166-1 Numeric format
	Countries []string `json:"countries[],omitempty"`
}

// Transparency calls GET /transparency/procedures_contacts_single_window.
func (s *Service) Transparency(ctx context.Context, p TransparencyParams) (json.RawMessage, error) {
	q := url.Values{}
	transport.Add(q, "countries[]", p.Countries)
	return transport.Decode[json.RawMessage](ctx, s.client, transport.Request{
		Service: "tfad",
		Path:    "/transparency/procedures_contacts_single_window",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}
