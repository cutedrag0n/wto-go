package timeseries

// DataPointExtended follows the WTO published schema.
type DataPointExtended struct {
	IndicatorCategoryCode             *string  `json:"indicatorCategoryCode,omitempty"`
	IndicatorCategory                 *string  `json:"indicatorCategory,omitempty"`
	IndicatorCode                     *string  `json:"indicatorCode,omitempty"`
	Indicator                         *string  `json:"indicator,omitempty"`
	ReportingEconomyCode              *string  `json:"reportingEconomyCode,omitempty"`
	ReportingEconomy                  *string  `json:"reportingEconomy,omitempty"`
	PartnerEconomyCode                *string  `json:"partnerEconomyCode,omitempty"`
	PartnerEconomy                    *string  `json:"partnerEconomy,omitempty"`
	ProductOrSectorClassificationCode *string  `json:"productOrSectorClassificationCode,omitempty"`
	ProductOrSectorClassification     *string  `json:"productOrSectorClassification,omitempty"`
	ProductOrSectorCode               *string  `json:"productOrSectorCode,omitempty"`
	ProductOrSector                   *string  `json:"productOrSector,omitempty"`
	PeriodCode                        *string  `json:"periodCode,omitempty"`
	Period                            *string  `json:"period,omitempty"`
	FrequencyCode                     *string  `json:"frequencyCode,omitempty"`
	Frequency                         *string  `json:"frequency,omitempty"`
	UnitCode                          *string  `json:"unitCode,omitempty"`
	Unit                              *string  `json:"unit,omitempty"`
	Year                              *int     `json:"year,omitempty"`
	ValueFlagCode                     *string  `json:"valueFlagCode,omitempty"`
	ValueFlag                         *string  `json:"valueFlag,omitempty"`
	TextValue                         *string  `json:"textValue,omitempty"`
	Value                             *float64 `json:"value,omitempty"`
}

// Metadata follows the WTO published schema.
type Metadata struct {
	MetadataCategoryCode              *string  `json:"metadataCategoryCode,omitempty"`
	IndicatorCategoryCode             *string  `json:"indicatorCategoryCode,omitempty"`
	IndicatorCode                     *string  `json:"indicatorCode,omitempty"`
	ReportingEconomyCode              *string  `json:"reportingEconomyCode,omitempty"`
	PartnerEconomyCode                *string  `json:"partnerEconomyCode,omitempty"`
	ProductOrSectorClassificationCode *string  `json:"productOrSectorClassificationCode,omitempty"`
	ProductOrSectorCode               *string  `json:"productOrSectorCode,omitempty"`
	PeriodCode                        *string  `json:"periodCode,omitempty"`
	FrequencyCode                     *string  `json:"frequencyCode,omitempty"`
	UnitCode                          *string  `json:"unitCode,omitempty"`
	Value                             *float64 `json:"value,omitempty"`
}

// Topic follows the WTO published schema.
type Topic struct {
	ID        int     `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	SortOrder int     `json:"sortOrder,omitempty"`
}

// Frequency follows the WTO published schema.
type Frequency struct {
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Period follows the WTO published schema.
type Period struct {
	Code          *string `json:"code,omitempty"`
	Name          *string `json:"name,omitempty"`
	Description   *string `json:"description,omitempty"`
	FrequencyCode *string `json:"frequencyCode,omitempty"`
	DisplayOrder  int     `json:"displayOrder,omitempty"`
}

// Unit follows the WTO published schema.
type Unit struct {
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
}

// IndicatorCategory follows the WTO published schema.
type IndicatorCategory struct {
	Code       *string `json:"code,omitempty"`
	Name       *string `json:"name,omitempty"`
	ParentCode *string `json:"parentCode,omitempty"`
	SortOrder  int     `json:"sortOrder,omitempty"`
}

// Indicator follows the WTO published schema.
type Indicator struct {
	Code                             *string `json:"code,omitempty"`
	Name                             *string `json:"name,omitempty"`
	CategoryCode                     *string `json:"categoryCode,omitempty"`
	CategoryLabel                    *string `json:"categoryLabel,omitempty"`
	SubcategoryCode                  *string `json:"subcategoryCode,omitempty"`
	SubcategoryLabel                 *string `json:"subcategoryLabel,omitempty"`
	UnitCode                         *string `json:"unitCode,omitempty"`
	UnitLabel                        *string `json:"unitLabel,omitempty"`
	StartYear                        int     `json:"startYear,omitempty"`
	EndYear                          int     `json:"endYear,omitempty"`
	FrequencyCode                    *string `json:"frequencyCode,omitempty"`
	FrequencyLabel                   *string `json:"frequencyLabel,omitempty"`
	NumberReporters                  int     `json:"numberReporters,omitempty"`
	NumberPartners                   *int    `json:"numberPartners,omitempty"`
	ProductSectorClassificationCode  *string `json:"productSectorClassificationCode,omitempty"`
	ProductSectorClassificationLabel *string `json:"productSectorClassificationLabel,omitempty"`
	HasMetadata                      *string `json:"hasMetadata,omitempty"`
	NumberDecimals                   int     `json:"numberDecimals,omitempty"`
	NumberDatapoints                 int     `json:"numberDatapoints,omitempty"`
	UpdateFrequency                  *string `json:"updateFrequency,omitempty"`
	Description                      *string `json:"description,omitempty"`
	SortOrder                        int     `json:"sortOrder,omitempty"`
}

// TerritoryGroup follows the WTO published schema.
type TerritoryGroup struct {
	Code         *string `json:"code,omitempty"`
	Name         *string `json:"name,omitempty"`
	DisplayOrder int     `json:"displayOrder,omitempty"`
}

// Territory follows the WTO published schema.
type Territory struct {
	Code         *string `json:"code,omitempty"`
	ISO3A        *string `json:"iso3A,omitempty"`
	Name         *string `json:"name,omitempty"`
	DisplayOrder int     `json:"displayOrder,omitempty"`
}

// ProductClassification follows the WTO published schema.
type ProductClassification struct {
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Product follows the WTO published schema.
type Product struct {
	Code                  *string `json:"code,omitempty"`
	Name                  *string `json:"name,omitempty"`
	Note                  *string `json:"note,omitempty"`
	ProductClassification *string `json:"productClassification,omitempty"`
	CodeUnique            *string `json:"codeUnique,omitempty"`
	DisplayOrder          *int    `json:"displayOrder,omitempty"`
	Hierarchy             *string `json:"hierarchy,omitempty"`
}

// Years follows the WTO published schema.
type Years struct {
	Year int `json:"year,omitempty"`
}

// ValueFlag follows the WTO published schema.
type ValueFlag struct {
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}
