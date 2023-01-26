package models

type DisableCardResponse struct {
	Card   Card    `json:"card"`
	Errors []Error `json:"errors"`
}
