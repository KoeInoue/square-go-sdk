package models

// Customer represents a Square customer profile in the Customer Directory of a Square seller
type Customer struct {
	ID             string              `json:"id"`
	CreatedAt      string              `json:"created_at"`
	UpdatedAt      string              `json:"updated_at"`
	Cards          []Card              `json:"cards"`
	GivenName      string              `json:"given_name"`
	FamilyName     string              `json:"family_name"`
	Nickname       string              `json:"nickname"`
	CompanyName    string              `json:"company_name"`
	EmailAddress   string              `json:"email_address"`
	Address        Address             `json:"address"`
	PhoneNumber    string              `json:"phone_number"`
	Birthday       string              `json:"birthday"`
	ReferenceId    string              `json:"reference_id"`
	Note           string              `json:"note"`
	Preferences    CustomerPreferences `json:"preferences"`
	CreationSource string              `json:"creation_source"`
	GroupIds       []string            `json:"group_ids"`
	SegmentIds     []string            `json:"segment_ids"`
	Version        int64               `json:"version"`
	TaxIds         CustomerTaxIds      `json:"tax_ids"`
}
