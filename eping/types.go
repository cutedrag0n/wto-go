package eping

import "encoding/json"

// SearchResult follows the WTO published schema.
type SearchResult struct {
	Items       []Notification `json:"items,omitempty"`
	CurrentPage *int           `json:"currentPage,omitempty"`
	PageSize    *int           `json:"pageSize,omitempty"`
	TotalCount  *int           `json:"totalCount,omitempty"`
}

// Notification follows the WTO published schema.
type Notification struct {
	ID                                           *string           `json:"id,omitempty"`
	Area                                         *string           `json:"area,omitempty"`
	NotificationType                             *string           `json:"notificationType,omitempty"`
	DistributionDate                             *string           `json:"distributionDate,omitempty"`
	CommentDeadlineDate                          json.RawMessage   `json:"commentDeadlineDate,omitempty"`
	DocumentSymbol                               *string           `json:"documentSymbol,omitempty"`
	DocumentSymbolOrdered                        *string           `json:"documentSymbolOrdered,omitempty"`
	NotifyingMember                              *string           `json:"notifyingMember,omitempty"`
	NotifyingMemberCode                          *string           `json:"notifyingMemberCode,omitempty"`
	Title                                        *string           `json:"title,omitempty"`
	Description                                  *string           `json:"description,omitempty"`
	ProductsFreeText                             *string           `json:"productsFreeText,omitempty"`
	HSCodes                                      []json.RawMessage `json:"hsCodes,omitempty"`
	ICSCodes                                     []ICSCode         `json:"icsCodes,omitempty"`
	Keywords                                     []json.RawMessage `json:"keywords,omitempty"`
	SPSKeywords                                  []json.RawMessage `json:"spsKeywords,omitempty"`
	Objectives                                   []Objective       `json:"objectives,omitempty"`
	HSCodeText                                   *string           `json:"hsCodeText,omitempty"`
	ICSCodeText                                  *string           `json:"icsCodeText,omitempty"`
	CountriesAffected                            []json.RawMessage `json:"countriesAffected,omitempty"`
	SpecificRegionsOrCountries                   json.RawMessage   `json:"specificRegionsOrCountries,omitempty"`
	AllTradingPartners                           json.RawMessage   `json:"allTradingPartners,omitempty"`
	CodexAlimentariusCommision                   json.RawMessage   `json:"codexAlimentariusCommision,omitempty"`
	WorldOrganizationForAnimalHealth             json.RawMessage   `json:"worldOrganizationForAnimalHealth,omitempty"`
	InternationalPlantProtectionConvention       json.RawMessage   `json:"internationalPlantProtectionConvention,omitempty"`
	NotRelevant                                  json.RawMessage   `json:"notRelevant,omitempty"`
	ConformStandardYes                           json.RawMessage   `json:"conformStandardYes,omitempty"`
	ConformStandardNo                            json.RawMessage   `json:"conformStandardNo,omitempty"`
	ConformStandardNA                            json.RawMessage   `json:"conformStandardNA,omitempty"`
	TradeFacilitatingMeasure                     json.RawMessage   `json:"tradeFacilitatingMeasure,omitempty"`
	TechnicalRegulation                          *bool             `json:"technicalRegulation,omitempty"`
	TechnicalRegulationUrgent                    *bool             `json:"technicalRegulationUrgent,omitempty"`
	ConformityAssessmentProcedure                *bool             `json:"conformityAssessmentProcedure,omitempty"`
	ConformityAssessmentProcedureUrgent          *bool             `json:"conformityAssessmentProcedureUrgent,omitempty"`
	TechnicalRegulationLocalGovernment           *bool             `json:"technicalRegulationLocalGovernment,omitempty"`
	ConformityAssessmentProcedureLocalGovernment *bool             `json:"conformityAssessmentProcedureLocalGovernment,omitempty"`
	ArticleOther                                 *bool             `json:"articleOther,omitempty"`
	DolLink                                      *string           `json:"dolLink,omitempty"`
	NotifiedDocumentLink                         *string           `json:"notifiedDocumentLink,omitempty"`
	NotificationFormatID                         *string           `json:"notificationFormatId,omitempty"`
	Version                                      json.RawMessage   `json:"version,omitempty"`
	AffectedRegionsOrCountries                   json.RawMessage   `json:"affectedRegionsOrCountries,omitempty"`
	ProposedAdoptionDate                         *string           `json:"proposedAdoptionDate,omitempty"`
	ProposedEntryIntoForceDate                   *string           `json:"proposedEntryIntoForceDate,omitempty"`
	PublicationDate                              *string           `json:"publicationDate,omitempty"`
	RelevantDocumentTitle                        *string           `json:"relevantDocumentTitle,omitempty"`
	ConformStandardRecommendation                *string           `json:"conformStandardRecommendation,omitempty"`
	ProposedAdoptionDateText                     *string           `json:"proposedAdoptionDateText,omitempty"`
	ProposedAdoptionDateToBeDetermined           *bool             `json:"proposedAdoptionDateToBeDetermined,omitempty"`
	ProposedEntryIntoForceDateText               *string           `json:"proposedEntryIntoForceDateText,omitempty"`
	ProposedEntryIntoForceFromAdoption6Months    *bool             `json:"proposedEntryIntoForceFromAdoption6Months,omitempty"`
	ProposedEntryIntoForceDateToBeDetermined     *bool             `json:"proposedEntryIntoForceDateToBeDetermined,omitempty"`
}

// ICSCode follows the WTO published schema.
type ICSCode struct {
	ID   *int    `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Objective follows the WTO published schema.
type Objective struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
