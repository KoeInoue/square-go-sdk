package models

// Represents the tax ID associated with a [customer profile]($m/Customer). The corresponding `tax_ids` field is available only for customers of sellers in EU countries or the United Kingdom.
type CustomerTaxIds struct {
	EuVat string `json:"eu_vat"`
}
