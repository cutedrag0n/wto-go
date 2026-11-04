package timeseries

import (
	"context"
	"net/url"

	"github.com/cutedrag0n/wto-go/internal/transport"
)

// TopicsParams configures Topics. Omitted fields retain server defaults.
type TopicsParams struct {
	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Topics calls GET /topics.
func (s *Service) Topics(ctx context.Context, p TopicsParams) ([]Topic, error) {
	q := url.Values{}
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]Topic](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/topics",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// FrequenciesParams configures Frequencies. Omitted fields retain server defaults.
type FrequenciesParams struct {
	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Frequencies calls GET /frequencies.
func (s *Service) Frequencies(ctx context.Context, p FrequenciesParams) ([]Frequency, error) {
	q := url.Values{}
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]Frequency](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/frequencies",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// PeriodsParams configures Periods. Omitted fields retain server defaults.
type PeriodsParams struct {
	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Periods calls GET /periods.
func (s *Service) Periods(ctx context.Context, p PeriodsParams) ([]Period, error) {
	q := url.Values{}
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]Period](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/periods",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// UnitsParams configures Units. Omitted fields retain server defaults.
type UnitsParams struct {
	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Units calls GET /units.
func (s *Service) Units(ctx context.Context, p UnitsParams) ([]Unit, error) {
	q := url.Values{}
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]Unit](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/units",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// IndicatorCategoriesParams configures IndicatorCategories. Omitted fields retain server defaults.
type IndicatorCategoriesParams struct {
	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// IndicatorCategories calls GET /indicator_categories.
func (s *Service) IndicatorCategories(ctx context.Context, p IndicatorCategoriesParams) ([]IndicatorCategory, error) {
	q := url.Values{}
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]IndicatorCategory](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/indicator_categories",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// IndicatorsParams configures Indicators. Omitted fields retain server defaults.
type IndicatorsParams struct {
	// Indicators: Indicator : all Indicator Code : Filter on one specific indicator
	Indicators string `json:"i,omitempty"`

	// Name: Indicator name (or part of the name)
	Name string `json:"name,omitempty"`

	// Topics: Topics (comma separated ids). See Topics.
	Topics string `json:"t,omitempty"`

	// ProductClassifications: Product Classifications : all : all classifications none : no product dimension comma separated codes : See Classifications.
	ProductClassifications string `json:"pc,omitempty"`

	// TradePartners: Trade partner : yes no all
	TradePartners string `json:"tp,omitempty"`

	// Frequency: Frequency : A : annualy S : semiannualy Q : quarterly M : monthly all
	Frequency string `json:"frq,omitempty"`

	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Indicators calls GET /indicators.
func (s *Service) Indicators(ctx context.Context, p IndicatorsParams) ([]Indicator, error) {
	q := url.Values{}
	transport.Add(q, "i", p.Indicators)
	transport.Add(q, "name", p.Name)
	transport.Add(q, "t", p.Topics)
	transport.Add(q, "pc", p.ProductClassifications)
	transport.Add(q, "tp", p.TradePartners)
	transport.Add(q, "frq", p.Frequency)
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]Indicator](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/indicators",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// RegionsParams configures Regions. Omitted fields retain server defaults.
type RegionsParams struct {
	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Regions calls GET /territory/regions.
func (s *Service) Regions(ctx context.Context, p RegionsParams) ([]TerritoryGroup, error) {
	q := url.Values{}
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]TerritoryGroup](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/territory/regions",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// GroupsParams configures Groups. Omitted fields retain server defaults.
type GroupsParams struct {
	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Groups calls GET /territory/groups.
func (s *Service) Groups(ctx context.Context, p GroupsParams) ([]TerritoryGroup, error) {
	q := url.Values{}
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]TerritoryGroup](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/territory/groups",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// ReportersParams configures Reporters. Omitted fields retain server defaults.
type ReportersParams struct {
	// Name: Reporter name (or part of name)
	Name string `json:"name,omitempty"`

	// IndividualGroups: Individual / Group economies : individual group all
	IndividualGroups string `json:"ig,omitempty"`

	// Regions: Regions (comma separated codes). See Geographical regions.
	Regions string `json:"reg,omitempty"`

	// Groups: Groups (comma separated codes). See Economic groups.
	Groups string `json:"gp,omitempty"`

	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Reporters calls GET /reporters.
func (s *Service) Reporters(ctx context.Context, p ReportersParams) ([]Territory, error) {
	q := url.Values{}
	transport.Add(q, "name", p.Name)
	transport.Add(q, "ig", p.IndividualGroups)
	transport.Add(q, "reg", p.Regions)
	transport.Add(q, "gp", p.Groups)
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]Territory](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/reporters",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// PartnersParams configures Partners. Omitted fields retain server defaults.
type PartnersParams struct {
	// Name: Reporter name (or part of name)
	Name string `json:"name,omitempty"`

	// IndividualGroups: Individual / Group economies : individual group all
	IndividualGroups string `json:"ig,omitempty"`

	// Regions: Regions (comma separated codes). See Geographical regions.
	Regions string `json:"reg,omitempty"`

	// Groups: Groups (comma separated codes). See Economic groups.
	Groups string `json:"gp,omitempty"`

	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Partners calls GET /partners.
func (s *Service) Partners(ctx context.Context, p PartnersParams) ([]Territory, error) {
	q := url.Values{}
	transport.Add(q, "name", p.Name)
	transport.Add(q, "ig", p.IndividualGroups)
	transport.Add(q, "reg", p.Regions)
	transport.Add(q, "gp", p.Groups)
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]Territory](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/partners",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// ClassificationsParams configures Classifications. Omitted fields retain server defaults.
type ClassificationsParams struct {
	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Classifications calls GET /product_classifications.
func (s *Service) Classifications(ctx context.Context, p ClassificationsParams) ([]ProductClassification, error) {
	q := url.Values{}
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]ProductClassification](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/product_classifications",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// ProductsParams configures Products. Omitted fields retain server defaults.
type ProductsParams struct {
	// Name: Product/sector name (or part of name)
	Name string `json:"name,omitempty"`

	// Products: Product/sector Classifications (comma separated codes). See Classifications.
	ProductClassifications string `json:"pc,omitempty"`

	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// Products calls GET /products.
func (s *Service) Products(ctx context.Context, p ProductsParams) ([]Product, error) {
	q := url.Values{}
	transport.Add(q, "name", p.Name)
	transport.Add(q, "pc", p.ProductClassifications)
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]Product](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/products",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// YearsParams configures Years. Omitted fields retain server defaults.
type YearsParams struct {
}

// Years calls GET /years.
func (s *Service) Years(ctx context.Context, p YearsParams) ([]Years, error) {
	q := url.Values{}
	return transport.Decode[[]Years](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/years",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}

// ValueFlagsParams configures ValueFlags. Omitted fields retain server defaults.
type ValueFlagsParams struct {
	// Language: Language id : 1 : English 2 : French 3 : Spanish
	Language *int `json:"lang,omitempty"`
}

// ValueFlags calls GET /value_flags.
func (s *Service) ValueFlags(ctx context.Context, p ValueFlagsParams) ([]ValueFlag, error) {
	q := url.Values{}
	transport.Add(q, "lang", p.Language)
	return transport.Decode[[]ValueFlag](ctx, s.client, transport.Request{
		Service: "timeseries",
		Path:    "/value_flags",
		Method:  "GET",
		Query:   q,
		Data:    false,
	})
}
