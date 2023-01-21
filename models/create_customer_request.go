package models

// CreateCustomerRequest defines the body parameters that
// can be included in a request to the `CreateCustomer` endpoint.
type CreateCustomerRequest struct {
	IdempotencyKey string         `json:"idempotency_key"`
	GivenName      string         `json:"given_name,omitempty"`
	FamilyName     string         `json:"family_name,omitempty"`
	CompanyName    string         `json:"company_name,omitempty"`
	Nickname       string         `json:"nickname,omitempty"`
	EmailAddress   string         `json:"email_address,omitempty"`
	Address        Address        `json:"address,omitempty"`
	PhoneNumber    string         `json:"phone_number,omitempty"`
	ReferenceId    string         `json:"reference_id"`
	Note           string         `json:"note,omitempty"`
	Birthday       string         `json:"birthday,omitempty"`
	TaxIds         CustomerTaxIds `json:"tax_ids,omitempty"`
}
