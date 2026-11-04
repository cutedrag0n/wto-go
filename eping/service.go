// Package eping provides the WTO eping API. Construct its Service through wto.NewClient.
package eping

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/cutedrag0n/wto-go/internal/transport"
)

// Service uses a shared authenticated transport.
type Service struct{ client *transport.Client }

// NewService is used by the root client to wire this service.
func NewService(c *transport.Client) *Service { return &Service{client: c} }

// SearchNotificationsParams configures SearchNotifications. Omitted fields retain server defaults.
type SearchNotificationsParams struct {
	// Language: Language of the notifications you are requesting (1 - English, 2 - French, 3 - Spanish)
	Language string `json:"language,omitempty"`

	// DomainIDs: Domain of the notification you want to retrieve, either 1 - TBT, 2 - SPS. Not specifying the domainIds will return notifications for both TBT and SPS
	DomainIDs *int `json:"domainIds,omitempty"`

	// DocumentSymbol: Search for a particular document symbol. Specifying "USA/1234" (without the quotes) will return all notifications where the document symbol contains "USA/1234" (e.g. G/TBT/N/USA/1234, G/SPS/N/USA/1234, G/SPS/N/USA/1234/Add.1, etc.)
	DocumentSymbol string `json:"documentSymbol,omitempty"`

	// DistributionDateFrom: Search for notifications which are issued on or after the date specified in this parameter. yyyy-MM-dd format. Specifying 2022-08-01 will return notifications that were issued on or after 8 August 8, 2022.
	DistributionDateFrom string `json:"distributionDateFrom,omitempty"`

	// DistributionDateTo: Search for notifications which are issued on or before the date specified in this parameter. yyyy-MM-dd format. Specifying "2022-08-01" will return notifications that were issued on or before 8 August 8, 2022.
	DistributionDateTo string `json:"distributionDateTo,omitempty"`

	// CountryIDs: Search for notifications issued by a particular Member. See members API for list of Country Ids (use value field)
	CountryIDs []string `json:"countryIds,omitempty"`

	// HS: Search for notifications related to specified HS Codes. To search for multiple HS Codes, submit comma-separated values of HS Codes (e.g. 0102,0103,etc.). When combined with "ics" parameter, system performs an "OR" condition (e.g. &hs=01,02&ics=01.030,01.040 will return notifications related to HS 01 OR HS 02 OR ICS 01.030 OR ICS 01.040)
	HS string `json:"hs,omitempty"`

	// ICS: Search for notifications related to specified ICS Codes. To search for multiple ICS Codes, submit comma-separated values of ICS Codes (e.g. 0102,0103,etc.). When combined with "hs" parameter, system performs an "OR" condition (e.g. &hs=01,02&ics=01.030,01.040 will return notifications related to HS 01 OR HS 02 OR ICS 01.030 OR ICS 01.040)
	ICS string `json:"ics,omitempty"`

	// FreeText: Search for any term related to a notification.
	FreeText string `json:"freeText,omitempty"`

	// Page: Page number of the notifications you want to retrieve. Default value is 1.
	Page *int `json:"page,omitempty"`

	// PageSize: Number of notifications to be retrieved per call. Default value is 20. Maximum value is 100.
	PageSize *int `json:"pageSize,omitempty"`
}

// SearchNotifications calls GET /notifications/search.
func (s *Service) SearchNotifications(ctx context.Context, p SearchNotificationsParams) (SearchResult, error) {
	if p.Language == "" {
		var zero SearchResult
		return zero, fmt.Errorf("wto: language is required")
	}
	q := url.Values{}
	transport.Add(q, "language", p.Language)
	transport.Add(q, "domainIds", p.DomainIDs)
	transport.Add(q, "documentSymbol", p.DocumentSymbol)
	transport.Add(q, "distributionDateFrom", p.DistributionDateFrom)
	transport.Add(q, "distributionDateTo", p.DistributionDateTo)
	transport.Add(q, "countryIds", p.CountryIDs)
	transport.Add(q, "hs", p.HS)
	transport.Add(q, "ics", p.ICS)
	transport.Add(q, "freeText", p.FreeText)
	transport.Add(q, "page", p.Page)
	transport.Add(q, "pageSize", p.PageSize)
	return transport.Decode[SearchResult](ctx, s.client, transport.Request{
		Service: "eping",
		Path:    "/notifications/search",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// MembersParams configures Members. Omitted fields retain server defaults.
type MembersParams struct {
	// Language: Language of the notifications you are requesting (1 - English, 2 - French, 3 - Spanish)
	Language string `json:"language,omitempty"`
}

// Members calls GET /members.
func (s *Service) Members(ctx context.Context, p MembersParams) (json.RawMessage, error) {
	if p.Language == "" {
		var zero json.RawMessage
		return zero, fmt.Errorf("wto: language is required")
	}
	q := url.Values{}
	transport.Add(q, "language", p.Language)
	return transport.Decode[json.RawMessage](ctx, s.client, transport.Request{
		Service: "eping",
		Path:    "/members",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}
