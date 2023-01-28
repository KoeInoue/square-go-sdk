package models

type CancelSubscriptionResponse struct {
	Subscription       Subscription       `json:"subscription,omitempty"`
	SubscriptionAction SubscriptionAction `json:"subscription_action,omitempty"`
	Errors             []Error            `json:"errors,omitempty"`
}
