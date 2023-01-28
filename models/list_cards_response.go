package models

type ListCardsResponse struct {
	Cards  []Card  `json:"cards,omitempty"`
	Errors []Error `json:"errors,omitempty"`
	Cursor string  `json:"cursor,omitempty"`
}
