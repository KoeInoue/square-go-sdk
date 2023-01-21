package models

type Subscription struct {
	Id                 string               `json:"id,omitempty"`
	LocationId         string               `json:"location_id,omitempty"`
	PlanId             string               `json:"plan_id,omitempty"`
	CustomerId         string               `json:"customer_id,omitempty"`
	StartDate          string               `json:"start_date,omitempty"`
	CanceledDate       string               `json:"canceled_date,omitempty"`
	ChargedThroughDate string               `json:"charged_through_date,omitempty"`
	Status             string               `json:"status,omitempty"`
	TaxPercentage      string               `json:"tax_percentage,omitempty"`
	InvoiceIds         []string             `json:"invoice_ids,omitempty"`
	PriceOverrideMoney Money                `json:"price_override_money,omitempty"`
	Version            int64                `json:"version,omitempty"`
	CreatedAt          string               `json:"created_at,omitempty"`
	CardId             string               `json:"card_id,omitempty"`
	Timezone           string               `json:"timezone,omitempty"`
	Source             SubscriptionSource   `json:"source,omitempty"`
	Actions            []SubscriptionAction `json:"actions,omitempty"`
}
