package models

// CreateCustomerResponse defines the fields that are included in the response body of
// a request to the `CreateCustomer` endpoint.
// Either `errors` or `customer` is present in a given response (never both).
type CreateCustomerResponse struct {
	Customer Customer `json:"customer"`
	Errors   []Error  `json:"errors"`
}
