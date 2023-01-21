package models

type CatalogObject struct {
	Type                  string `json:"type"`
	Id                    string `json:"id"`
	UpdatedAt             string `json:"updated_at"`
	Version               int64  `json:"version"`
	IsDeleted             bool   `json:"is_deleted"`
	CustomAttributeValues string `json:"custom_attribute_values"`
	// CatalogV1Ids []CatalogV1Id // TODO: implement this
	PresentAtAllLocations bool     `json:"present_at_all_locations"`
	PresentAtLocationIds  []string `json:"present_at_location_ids"`
	AbsentAtLocationIds   []string `json:"absent_at_location_ids"`
	// TODO: implement below
	// ItemData CatalogItem
	// CategoryData CatalogCategory
	// ItemVariationData CatalogItemVariation
	// TaxData CatalogTax
	// DiscountData CatalogDiscount
	// ModifierListData CatalogModifierList
	// ModifierData CatalogModifier
	// TimePeriodData CatalogTimePeriod
	// ProductSetData CatalogProductSet
	// PricingRuleData CatalogPricingRule
	// ImageData CatalogImage
	// MeasurementUnitData CatalogMeasurementUnit
	// SubscriptionPlanData CatalogSubscriptionPlan
	// ItemOptionData CatalogItemOption
	// ItemOptionValueData CatalogItemOptionValue
	// CustomAttributeDefinitionData CatalogCustomAttributeDefinition
	// QuickAmountsSettingsData CatalogQuickAmountsSettings
}
