package models

// SearchSubscriptionsResponse represents the response body for searching subscriptions
type SearchSubscriptionsResponse struct {
	Errors        []Error        `json:"errors,omitempty"`
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
	Cursor        string         `json:"cursor,omitempty"`
}
