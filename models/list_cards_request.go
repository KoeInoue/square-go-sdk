package models

type ListCardsRequest struct {
	Cursor          string `json:"cursor,omitempty"`
	CustomerId      string `json:"customer_id,omitempty"`
	IncludeDisabled bool   `json:"include_disabled,omitempty"`
	ReferenceId     string `json:"reference_id,omitempty"`
	SortOrder       string `json:"sort_order,omitempty"`
}
