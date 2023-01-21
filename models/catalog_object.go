package models

type CatalogObject struct {
	Type                  string
	Id                    string
	UpdatedAt             string
	Version               int64
	IsDeleted             bool
	CustomAttributeValues string
	// CatalogV1Ids []CatalogV1Id // TODO: implement this
	PresentAtAllLocations bool
	PresentAtLocationIds  []string
	AbsentAtLocationIds   []string
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
