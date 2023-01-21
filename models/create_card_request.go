package models

// CreateCardRequest creates a card from the source (nonce, payment id, etc).t
type CreateCardRequest struct {
	IdempotencyKey    string `json:"idempotency_key"`
	SourceId          string `json:"source_id"`
	VerificationToken string `json:"verification_token,omitempty"`
	Card              Card   `json:"card"`
}
