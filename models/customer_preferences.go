package models

// CustomerPreferences represents communication preferences for the customer profile.
type CustomerPreferences struct {
	EmailUnsubscribed bool `json:"email_unsubscribed"`
}
