package models

// CreateSubscriptionRequest Defines input parameters in a request to the
// [CreateSubscription]($e/Subscriptions/CreateSubscription) endpoint.
type CreateSubscriptionRequest struct {
	CustomerId         string             `json:"customer_id"`
	LocationId         string             `json:"location_id"`
	PlanId             string             `json:"plan_id"`
	PlanVariationId    string             `json:"plan_variation_id"`
	IdempotencyKey     string             `json:"idempotency_key"`
	StartDate          string             `json:"start_date,omitempty"`
	CanceledDate       string             `json:"canceled_date,omitempty"`
	TaxPercentage      string             `json:"tax_percentage,omitempty"`
	PriceOverrideMoney Money              `json:"price_override_money,omitempty"`
	CardId             string             `json:"card_id,omitempty"`
	Timezone           string             `json:"timezone,omitempty"`
	Source             SubscriptionSource `json:"source,omitempty"`
}
