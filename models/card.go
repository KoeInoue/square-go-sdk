package models

// Card represents the payment details of a card to be used for payments.
type Card struct {
	ID             string  `json:"id,omitempty"`
	CardBrand      string  `json:"card_brand,omitempty"`
	Last4          string  `json:"last_4,omitempty"`
	ExpMonth       int64   `json:"exp_month,omitempty"`
	ExpYear        int64   `json:"exp_year,omitempty"`
	CardholderName string  `json:"cardholder_name,omitempty"`
	BillingAddress Address `json:"billing_address,omitempty"`
	Fingerprint    string  `json:"fingerprint,omitempty"`
	CustomerId     string  `json:"customer_id,omitempty"`
	MerchantId     string  `json:"merchant_id,omitempty"`
	ReferenceId    string  `json:"reference_id,omitempty"`
	Enabled        bool    `json:"enabled,omitempty"`
	CardType       string  `json:"card_type,omitempty"`
	PrepaidType    string  `json:"prepaid_type,omitempty"`
	Bin            string  `json:"bin,omitempty"`
	Version        int64   `json:"version,omitempty"`
	CardCoBrand    string  `json:"card_co_brand,omitempty"`
}
