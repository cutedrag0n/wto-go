package timeseries

import (
	"context"
	"fmt"
	"net/url"

	"github.com/cutedrag0n/wto-go/internal/transport"
)

// MetadataParams configures Metadata. Omitted fields retain server defaults.
type MetadataParams struct {
	// Indicators: Indicator code. See Indicators.
	Indicators string `json:"i,omitempty"`

	// Reporters: Reporting economies (comma separated codes). See Reporting economies.
	Reporters string `json:"r,omitempty"`

	// Partners: Partner economies where applicable (comma separated codes). See Partner economies.
	Partners string `json:"p,omitempty"`

	// Products: Products/sectors (comma separated codes) where applicable. See Products/sectors. default : includes all the product/sectors codes of the indicator's classification (except for HS : Include all the top level 2 digit products + the whole chapter 01) all : includes all the product/sectors codes of the indicator's classification HS2, HS4, HS6 : All HS codes at 2,4,6 digit-level AG,AGFOFI,MAIS,... : comma separated list of product codes
	Products string `json:"pc,omitempty"`

	// IncludeSubproducts: Include sub products/sectors. If true, all child items in the product/sector hierarchy are recursively included.
	IncludeSubproducts *bool `json:"spc,omitempty"`

	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Metadata calls GET /metadata.
func (s *Service) Metadata(ctx context.Context, p MetadataParams) ([]Metadata, error) {
	if p.Indicators == "" {
		var zero []Metadata
		return zero, fmt.Errorf("wto: i is required")
	}
	q := url.Values{}
	transport.Add(q, "i", p.Indicators)
	transport.Add(q, "r", p.Reporters)
	transport.Add(q, "p", p.Partners)
	transport.Add(q, "pc", p.Products)
	transport.Add(q, "spc", p.IncludeSubproducts)
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]Metadata](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/metadata",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// DataCountParams configures DataCount. Omitted fields retain server defaults.
type DataCountParams struct {
	// Indicators: Indicator code. See Indicators.
	Indicators string `json:"i,omitempty"`

	// Reporters: Reporting economies (comma separated codes). See Reporting economies.
	Reporters string `json:"r,omitempty"`

	// Partners: Partner economies where applicable (comma separated codes). See Partner economies.
	Partners string `json:"p,omitempty"`

	// Periods: Time period : default : The last eight years will be applied. all : All years are applied. YYYY : comma separated list of years. If the frequency of the requested indicator is monthly or quarterly, the resulting output will reflect this, e.g. 2010 will give an output of January-December or 2010Q1-Q4. YYYYQn : comma separated list of year + quarter. If the indicator frequency is monthly, the output will be given in months, e.g. 2010Q2 will produce an output of April, May, June. YYYYMM : comma separated list of year + month. e.g. 201007, 201507. YYYY-YYYY : range of years. e.g. 1995-2015. YYYYQn-YYYYQn : range of quarters. e.g. 2010Q1-2013Q3. YYYYMM-YYYYMM : range of months. e.g. 201501-201606.
	Periods string `json:"ps,omitempty"`

	// Products: Products/sectors (comma separated codes) where applicable. See Products/sectors. default : includes all the product/sector codes of the indicator's classification (except for HS : includes all the top level 2-digit products + the whole chapter 01) all : includes all the product/sector codes of the indicator's classification HS2, HS4, HS6 : All HS codes at 2,4,6-digit level AG,AGFOFI,MAIS,... : comma separated list of product codes
	Products string `json:"pc,omitempty"`

	// IncludeSubproducts: Include sub products/sectors.
	IncludeSubproducts *bool `json:"spc,omitempty"`
}

// DataCount calls GET /data_count.
func (s *Service) DataCount(ctx context.Context, p DataCountParams) (int, error) {
	if p.Indicators == "" {
		var zero int
		return zero, fmt.Errorf("wto: i is required")
	}
	q := url.Values{}
	transport.Add(q, "i", p.Indicators)
	transport.Add(q, "r", p.Reporters)
	transport.Add(q, "p", p.Partners)
	transport.Add(q, "ps", p.Periods)
	transport.Add(q, "pc", p.Products)
	transport.Add(q, "spc", p.IncludeSubproducts)
	return transport.Decode[int](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/data_count",
		Method:  "GET",
		Query:   q,
		Data:    true,
	})
}
