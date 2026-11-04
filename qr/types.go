package qr

// Product follows the WTO published schema.
type Product struct {
	ID          string         `json:"id,omitempty"`
	Code        string         `json:"code,omitempty"`
	Tier        Integer        `json:"tier,omitempty"`
	Leaf        Integer        `json:"leaf,omitempty"`
	Description TranslatedText `json:"description,omitempty"`
	HSVersion   string         `json:"hs_version,omitempty"`
}

// Member follows the WTO published schema.
type Member struct {
	Code string         `json:"code,omitempty"`
	Name TranslatedText `json:"name,omitempty"`
}

// Notification follows the WTO published schema.
type Notification struct {
	QRSerialNumber *int `json:"qr_sn,omitempty"`

	ReporterMember   Member   `json:"reporter_member,omitempty"`
	NotificationDt   string   `json:"notification_dt,omitempty"`
	DocumentSymbol   string   `json:"document_symbol,omitempty"`
	DocumentURL      string   `json:"document_url,omitempty"`
	OriginalLanguage string   `json:"original_language,omitempty"`
	Type             string   `json:"type,omitempty"`
	CoveredPeriods   []string `json:"covered_periods,omitempty"`
}

type QRID int

// TranslatedText follows the WTO published schema.
type TranslatedText struct {
	En string `json:"en,omitempty"`
	Fr string `json:"fr,omitempty"`
	Es string `json:"es,omitempty"`
}

// Measure follows the WTO published schema.
type Measure struct {
	MastCodes []string `json:"mast_codes,omitempty"`

	Flow        string         `json:"flow,omitempty"`
	Symbol      string         `json:"symbol,omitempty"`
	GroupName   string         `json:"group_name,omitempty"`
	Description TranslatedText `json:"description,omitempty"`
	Interpreted Boolean        `json:"interpreted,omitempty"`
}

// QR follows the WTO published schema.
type QR struct {
	// ReporterMember and Products reflect the current live list response; ReportingMember and AffectedProducts retain the published schema.
	ReporterMember Member              `json:"reporter_member,omitempty"`
	Products       []RestrictedProduct `json:"products,omitempty"`
	Details        string              `json:"details,omitempty"`

	ID                       QRID             `json:"id,omitempty"`
	InForceFrom              string           `json:"in_force_from,omitempty"`
	TerminationDt            string           `json:"termination_dt,omitempty"`
	ReportingMember          Member           `json:"reporting_member,omitempty"`
	GeneralDescription       string           `json:"general_description,omitempty"`
	Restrictions             []string         `json:"restrictions,omitempty"`
	NationalLegalBases       string           `json:"national_legal_bases,omitempty"`
	AdministrativeMechanisms string           `json:"administrative_mechanisms,omitempty"`
	Measures                 []Measure        `json:"measures,omitempty"`
	AffectedProducts         []Product        `json:"affected_products,omitempty"`
	ResponsibleEntities      []TranslatedText `json:"responsible_entities,omitempty"`
	ValidFrom                string           `json:"valid_from,omitempty"`
	ValidTo                  string           `json:"valid_to,omitempty"`
	VersionID                int              `json:"version_id,omitempty"`
	NotifiedIn               []Notification   `json:"notified_in,omitempty"`
}

// Page preserves the documented pagination envelope.
type Page[T any] struct {
	Data  []T       `json:"data"`
	Links PageLinks `json:"links"`
	Meta  PageMeta  `json:"meta"`
}
type PageLinks struct {
	First *string `json:"first"`
	Last  *string `json:"last"`
	Prev  *string `json:"prev"`
	Next  *string `json:"next"`
}
type PageMeta struct {
	CurrentPage int        `json:"current_page"`
	From        *int       `json:"from"`
	LastPage    int        `json:"last_page"`
	PerPage     int        `json:"per_page"`
	To          *int       `json:"to"`
	Total       int        `json:"total"`
	Path        string     `json:"path"`
	Links       []PageLink `json:"links"`
}
type PageLink struct {
	URL    *string `json:"url"`
	Label  string  `json:"label"`
	Active bool    `json:"active"`
}
type Detail struct {
	Data QR `json:"data"`
}
type HSVersion struct {
	Code    string `json:"code"`
	Version string `json:"version"`
	Label   string `json:"label"`
}

type RestrictedProduct struct {
	HSVersion   string `json:"hs_version"`
	ProductCode string `json:"product_code"`
}
