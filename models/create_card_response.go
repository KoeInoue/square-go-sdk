package models

// CreateCardResponse Defines the fields that are included in the response body of
// a request to the [CreateCard]($e/Cards/CreateCard) endpoint.
// Note: if there are errors processing the request, the card field will not be
// present.
type CreateCardResponse struct {
	Errors []Error `json:"errors,omitempty"`
	card   Card    `json:"card,omitempty"`
}
