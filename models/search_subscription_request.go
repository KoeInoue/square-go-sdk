package models

// SearchSubscriptionsRequest represents the request body for searching subscriptions
type SearchSubscriptionsRequest struct {
	Cursor  string                   `json:"cursor,omitempty"`
	Limit   int                      `json:"limit,omitempty"`
	Query   SearchSubscriptionsQuery `json:"query,omitempty"`
	Include []string                 `json:"include,omitempty"`
}

// SearchSubscriptionsQuery represents the query conditions for searching subscriptions
type SearchSubscriptionsQuery struct {
	// Define the fields as per the Square API documentation
	FilterConditions map[string]interface{} `json:"filter,omitempty"`
}
