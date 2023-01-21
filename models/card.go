package models

// Card represents the payment details of a card to be used for payments.
type Card struct {
	ID             string  `json:"id"`
	CardBrand      string  `json:"card_brand"`
	Last4          string  `json:"last_4"`
	ExpMonth       int64   `json:"exp_month"`
	ExpYear        int64   `json:"exp_year"`
	CardholderName string  `json:"cardholder_name"`
	BillingAddress Address `json:"billing_address"`
	Fingerprint    string  `json:"fingerprint"`
	CustomerId     string  `json:"customer_id"`
	MerchantId     string  `json:"merchant_id"`
	ReferenceId    string  `json:"reference_id"`
	Enabled        bool    `json:"enabled"`
	CardType       string  `json:"card_type"`
	PrepaidType    string  `json:"prepaid_type"`
	Bin            string  `json:"bin"`
	Version        int64   `json:"version"`
	CardCoBrand    string  `json:"card_co_brand"`
}
