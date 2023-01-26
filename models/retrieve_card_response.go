package models

type RetrieveCardResponse struct {
	Card   Card    `json:"card"`
	Errors []Error `json:"errors"`
}
