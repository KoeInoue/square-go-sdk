package models

// CreateCustomerRequest defines the body parameters that
// can be included in a request to the `CreateCustomer` endpoint.
type CreateCustomerRequest struct {
	IdempotencyKey string         `json:"idempotency_key"`
	GivenName      string         `json:"given_name"`
	FamilyName     string         `json:"family_name"`
	CompanyName    string         `json:"company_name"`
	Nickname       string         `json:"nickname"`
	EmailAddress   string         `json:"email_address"`
	Address        Address        `json:"address"`
	PhoneNumber    string         `json:"phone_number"`
	ReferenceId    string         `json:"reference_id"`
	Note           string         `json:"note"`
	Birthday       string         `json:"birthday"`
	TaxIds         CustomerTaxIds `json:"tax_ids"`
}
