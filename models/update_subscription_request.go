package models

type UpdateSubscriptionRequest struct {
	Subscription Subscription `json:"subscription,omitempty"`
}
