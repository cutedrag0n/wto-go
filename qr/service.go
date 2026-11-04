// Package qr provides the WTO qr API. Construct its Service through wto.NewClient.
package qr

import (
	"context"
	"fmt"
	"net/url"

	"github.com/cutedrag0n/wto-go/internal/transport"
)

// Service uses a shared authenticated transport.
type Service struct{ client *transport.Client }

// NewService is used by the root client to wire this service.
func NewService(c *transport.Client) *Service { return &Service{client: c} }

// MembersParams configures Members. Omitted fields retain server defaults.
type MembersParams struct {
	// MemberCode: WTO member code
	MemberCode string `json:"member_code,omitempty"`

	// Name: Member's name
	Name string `json:"name,omitempty"`

	// Page: Page number
	Page *int `json:"page,omitempty"`
}

// Members calls GET /members.
func (s *Service) Members(ctx context.Context, p MembersParams) (Page[Member], error) {
	q := url.Values{}
	transport.Add(q, "member_code", p.MemberCode)
	transport.Add(q, "name", p.Name)
	transport.Add(q, "page", p.Page)
	return transport.Decode[Page[Member]](ctx, s.client, transport.Request{
		Service: "qr",
		Path:    "/members",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// ProductsParams configures Products. Omitted fields retain server defaults.
type ProductsParams struct {
	// HSVersion: HS version (h1, h2, h3, h4, h5, h6, h7)
	HSVersion string `json:"hs_version,omitempty"`

	// Code: Product code to fully or partially (prefix) match in the results
	Code string `json:"code,omitempty"`

	// Page: Page number
	Page *int `json:"page,omitempty"`

	// Description: Search in the description of the products.
	Description string `json:"description,omitempty"`
}

// Products calls GET /products.
func (s *Service) Products(ctx context.Context, p ProductsParams) (Page[Product], error) {
	if p.HSVersion == "" {
		var zero Page[Product]
		return zero, fmt.Errorf("wto: hs_version is required")
	}
	q := url.Values{}
	transport.Add(q, "hs_version", p.HSVersion)
	transport.Add(q, "code", p.Code)
	transport.Add(q, "page", p.Page)
	transport.Add(q, "description", p.Description)
	return transport.Decode[Page[Product]](ctx, s.client, transport.Request{
		Service: "qr",
		Path:    "/products",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// NotificationsParams configures Notifications. Omitted fields retain server defaults.
type NotificationsParams struct {
	// Page: Page number
	Page *int `json:"page,omitempty"`

	// ReporterMemberCode: WTO member code of the notifying member
	ReporterMemberCode string `json:"reporter_member_code,omitempty"`

	// NotificationYear: Year of submission
	NotificationYear *int `json:"notification_year,omitempty"`
}

// Notifications calls GET /notifications.
func (s *Service) Notifications(ctx context.Context, p NotificationsParams) (Page[Notification], error) {
	q := url.Values{}
	transport.Add(q, "page", p.Page)
	transport.Add(q, "reporter_member_code", p.ReporterMemberCode)
	transport.Add(q, "notification_year", p.NotificationYear)
	return transport.Decode[Page[Notification]](ctx, s.client, transport.Request{
		Service: "qr",
		Path:    "/notifications",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// RestrictionParams configures Restriction. Omitted fields retain server defaults.
type RestrictionParams struct {
	// QRID: Id of the QR
	QRID *int `json:"qrId,omitempty"`
}

// Restriction calls GET /qrs/{qrId}.
func (s *Service) Restriction(ctx context.Context, p RestrictionParams) (Detail, error) {
	if p.QRID == nil {
		var zero Detail
		return zero, fmt.Errorf("wto: qrId is required")
	}
	q := url.Values{}
	return transport.Decode[Detail](ctx, s.client, transport.Request{
		Service: "qr",
		Path:    fmt.Sprintf("/qrs/%d", *p.QRID),
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// HSVersionsParams configures HSVersions. Omitted fields retain server defaults.
type HSVersionsParams struct {
}

// HSVersions calls GET /hs-versions.
func (s *Service) HSVersions(ctx context.Context, p HSVersionsParams) (Page[HSVersion], error) {
	q := url.Values{}
	return transport.Decode[Page[HSVersion]](ctx, s.client, transport.Request{
		Service: "qr",
		Path:    "/hs-versions",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// RestrictionsParams configures Restrictions. Omitted fields retain server defaults.
type RestrictionsParams struct {
	// Page: Page number
	Page *int `json:"page,omitempty"`

	// ReporterMemberCode: WTO country code of the reporting member
	ReporterMemberCode string `json:"reporter_member_code,omitempty"`

	// InForceOnly: If true, only QRs currently in force will be returned
	InForceOnly *bool `json:"in_force_only,omitempty"`

	// YearOfEntryIntoForce: Year when the QRs went into force
	YearOfEntryIntoForce *int `json:"year_of_entry_into_force,omitempty"`

	// ProductCodes: Comma separated list of product codes to fully or partially (prefix) match any of the provided codes in any HS version.
	ProductCodes string `json:"product_codes,omitempty"`
	// ProductIDs: Comma separated list of product ids (in {version code}-{6 digit code} format, ex. h7-940191). The result only contains QRs where any of the product ids have a partial match (ex.: h7-0104 will match Qrs having h7-0104, h7-010410 and h7-01042
	ProductIDs string `json:"product_ids,omitempty"`
}

// Restrictions calls GET /qrs.
func (s *Service) Restrictions(ctx context.Context, p RestrictionsParams) (Page[QR], error) {
	q := url.Values{}
	transport.Add(q, "page", p.Page)
	transport.Add(q, "reporter_member_code", p.ReporterMemberCode)
	transport.Add(q, "in_force_only", p.InForceOnly)
	transport.Add(q, "year_of_entry_into_force", p.YearOfEntryIntoForce)
	transport.Add(q, "product_codes", p.ProductCodes)
	transport.Add(q, "product_ids", p.ProductIDs)
	return transport.Decode[Page[QR]](ctx, s.client, transport.Request{
		Service: "qr",
		Path:    "/qrs",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}
