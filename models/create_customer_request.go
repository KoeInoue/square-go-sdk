package models

type CreateCustomerRequest struct {
    EmailAddress   string `json:"email_address"`
    IdempotencyKey string `json:"idempotency_key"`
    Nickname       string `json:"nickname"`
}
