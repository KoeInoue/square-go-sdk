package models

type CatalogResponse struct {
	Errors []Error       `json:"errors,omitempty"`
	Cursor string        `json:"cursor,omitempty"`
	Object CatalogObject `json:"object,omitempty"`
}

type CatalogObject struct {
	Type                  string               `json:"type"` // Correctly matches "SUBSCRIPTION_PLAN" or other types as per the JSON
	ID                    string               `json:"id"`
	UpdatedAt             string               `json:"updated_at"`
	CreatedAt             string               `json:"created_at"`
	Version               int64                `json:"version"`
	IsDeleted             bool                 `json:"is_deleted"`
	PresentAtAllLocations bool                 `json:"present_at_all_locations"`
	SubscriptionPlanData  SubscriptionPlanData `json:"subscription_plan_data,omitempty"`
	ItemData              SubscriptionPlanData `json:"item_data,omitempty"`
}

type SubscriptionPlanData struct {
	Name                       string                      `json:"name"`
	SubscriptionPlanVariations []SubscriptionPlanVariation `json:"subscription_plan_variations,omitempty"`
	Variations                 []SubscriptionPlanVariation `json:"variations,omitempty"`
	EligibleItemIds            []string                    `json:"eligible_item_ids"`
	AllItems                   bool                        `json:"all_items"`
}

type SubscriptionPlanVariation struct {
	Type                          string                        `json:"type"` // "SUBSCRIPTION_PLAN_VARIATION"
	ID                            string                        `json:"id"`
	UpdatedAt                     string                        `json:"updated_at"`
	CreatedAt                     string                        `json:"created_at"`
	Version                       int64                         `json:"version"`
	IsDeleted                     bool                          `json:"is_deleted"`
	PresentAtAllLocations         bool                          `json:"present_at_all_locations"`
	SubscriptionPlanVariationData SubscriptionPlanVariationData `json:"subscription_plan_variation_data"`
}

type SubscriptionPlanVariationData struct {
	Name               string  `json:"name"`
	Phases             []Phase `json:"phases"`
	SubscriptionPlanID string  `json:"subscription_plan_id"`
}

type Phase struct {
	UID     string  `json:"uid"`
	Cadence string  `json:"cadence"`
	Ordinal int     `json:"ordinal"`
	Pricing Pricing `json:"pricing"`
}

type Pricing struct {
	Type string `json:"type"` // "RELATIVE" or other types as per the JSON
}
