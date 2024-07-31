package models

type CatalogResponse struct {
	Errors  []Error         `json:"errors,omitempty"`
	Cursor  string          `json:"cursor,omitempty"`
	Objects []CatalogObject `json:"objects,omitempty"`
}

// CatalogObject defines the structure for the catalog object response
type CatalogObject struct {
	Type                  string               `json:"type"`
	ID                    string               `json:"id"`
	UpdatedAt             string               `json:"updated_at"`
	CreatedAt             string               `json:"created_at"`
	Version               int64                `json:"version"`
	IsDeleted             bool                 `json:"is_deleted"`
	PresentAtAllLocations bool                 `json:"present_at_all_locations"`
	SubscriptionPlanData  SubscriptionPlanData `json:"subscription_plan_data"`
}

// SubscriptionPlanData defines the structure for subscription plan details
type SubscriptionPlanData struct {
	Name                       string                      `json:"name"`
	SubscriptionPlanVariations []SubscriptionPlanVariation `json:"subscription_plan_variations"`
	EligibleItemIds            []string                    `json:"eligible_item_ids"`
	AllItems                   bool                        `json:"all_items"`
}

// SubscriptionPlanVariation defines the structure for variations in a subscription plan
type SubscriptionPlanVariation struct {
	Type                          string                        `json:"type"`
	ID                            string                        `json:"id"`
	UpdatedAt                     string                        `json:"updated_at"`
	CreatedAt                     string                        `json:"created_at"`
	Version                       int64                         `json:"version"`
	IsDeleted                     bool                          `json:"is_deleted"`
	PresentAtAllLocations         bool                          `json:"present_at_all_locations"`
	SubscriptionPlanVariationData SubscriptionPlanVariationData `json:"subscription_plan_variation_data"`
}

// SubscriptionPlanVariationData defines the data structure for a variation of a subscription plan
type SubscriptionPlanVariationData struct {
	Name               string  `json:"name"`
	Phases             []Phase `json:"phases"`
	SubscriptionPlanID string  `json:"subscription_plan_id"`
}

// Phase defines the specific phases of a subscription variation
type Phase struct {
	UID     string  `json:"uid"`
	Cadence string  `json:"cadence"`
	Ordinal int     `json:"ordinal"`
	Pricing Pricing `json:"pricing"`
}

// Pricing defines the pricing details for a phase
type Pricing struct {
	Type string `json:"type"`
}
