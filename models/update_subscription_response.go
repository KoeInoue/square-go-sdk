package models

type UpdateSubscriptionResponse struct {
	Subscription Subscription `json:"subscription,omitempty"`
	Errors       []Error      `json:"errors,omitempty"`
}
