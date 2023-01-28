package models

// DeleteCustomerResponse defines the fields that are included in the response body of
// a request to the `DeleteCustomer` endpoint.
type DeleteCustomerResponse struct {
	Errors []Error `json:"errors"`
}
