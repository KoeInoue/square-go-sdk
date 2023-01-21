package models

// CreateSubscriptionRequest Defines input parameters in a request to the
// [CreateSubscription]($e/Subscriptions/CreateSubscription) endpoint.
type CreateSubscriptionRequest struct {
	IdempotencyKey     string             `json:"idempotency_key"`
	LocationId         string             `json:"location_id"`
	PlanId             string             `json:"plan_id"`
	CustomerId         string             `json:"customer_id"`
	StartDate          string             `json:"start_date"`
	CanceledDate       string             `json:"canceled_date,omitempty"`
	TaxPercentage      string             `json:"tax_percentage,omitempty"`
	PriceOverrideMoney Money              `json:"price_override_money,omitempty"`
	CardId             string             `json:"card_id,omitempty"`
	Timezone           string             `json:"timezone,omitempty"`
	Source             SubscriptionSource `json:"source,omitempty"`
}
