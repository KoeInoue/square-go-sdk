package models

type CatalogObject struct {
	Type                  string   `json:"type"`
	ID                    string   `json:"id"`
	UpdatedAt             string   `json:"updated_at"`
	CreatedAt             string   `json:"created_at"`
	Version               int64    `json:"version"`
	IsDeleted             bool     `json:"is_deleted"`
	PresentAtAllLocations bool     `json:"present_at_all_locations"`
	ItemData              ItemData `json:"item_data"`
}

type ItemData struct {
	Name               string      `json:"name"`
	IsTaxable          bool        `json:"is_taxable"`
	Visibility         string      `json:"visibility"`
	Variations         []Variation `json:"variations"`
	ProductType        string      `json:"product_type"`
	SkipModifierScreen bool        `json:"skip_modifier_screen"`
	EcomVisibility     string      `json:"ecom_visibility"`
	IsArchived         bool        `json:"is_archived"`
}

type Variation struct {
	Type                  string            `json:"type"`
	ID                    string            `json:"id"`
	UpdatedAt             string            `json:"updated_at"`
	CreatedAt             string            `json:"created_at"`
	Version               int64             `json:"version"`
	IsDeleted             bool              `json:"is_deleted"`
	PresentAtAllLocations bool              `json:"present_at_all_locations"`
	ItemVariationData     ItemVariationData `json:"item_variation_data"`
}

type ItemVariationData struct {
	ItemID              string     `json:"item_id"`
	Name                string     `json:"name"`
	Ordinal             int        `json:"ordinal"`
	PricingType         string     `json:"pricing_type"`
	PriceMoney          PriceMoney `json:"price_money"`
	Sellable            bool       `json:"sellable"`
	Stockable           bool       `json:"stockable"`
	SubscriptionPlanIds []string   `json:"subscription_plan_ids"`
}

type PriceMoney struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
