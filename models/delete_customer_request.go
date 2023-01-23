package models

// DeleteCustomerRequest represents the request body for DeleteCustomer API
type DeleteCustomerRequest struct {
	CustomerId string `json:"customer_id"`
	Version    int64  `json:"version"`
}
