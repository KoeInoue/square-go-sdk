package models

// CreateCardRequest creates a card from the source (nonce, payment id, etc).t
type CreateCardRequest struct {
    IdempotencyKey string
    SourceId string
    VerificationToken string
    Card Card
}
