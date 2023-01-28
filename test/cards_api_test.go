package test

import (
	"testing"

	"github.com/KoeInoue/square-go-sdk/models"
	"github.com/google/uuid"
)

// TestCardsApi is a test for CardApi TODO: Add more test cases
func TestCardsApi(t *testing.T) {
	cardRes, err := client.CardApi.CreateCard(models.CreateCardRequest{
		SourceId:       "cnon:card-nonce-ok",
		IdempotencyKey: uuid.New().String(),
		Card: models.Card{
			CustomerId: testCustomer.ID,
		},
	})
	if err != nil {
		t.Error(err)
	}
	if len(cardRes.Errors) > 0 {
		t.Errorf("Square Error: %#v, Code: %s", cardRes.Errors[0].Detail, cardRes.Errors[0].Code)
	}

	retriedCardRes, err := client.CardApi.RetrieveCard(cardRes.Card.ID)
	if err != nil {
		t.Error(err)
	}

	if len(retriedCardRes.Errors) > 0 {
		t.Errorf("Square Error: %#v, Code: %s", retriedCardRes.Errors[0].Detail, retriedCardRes.Errors[0].Code)
	}

	disCardRes, err := client.CardApi.DisableCard(cardRes.Card.ID)
	if err != nil {
		t.Error(err)
	}

	if len(disCardRes.Errors) > 0 {
		t.Errorf("Square Error: %#v, Code: %s", disCardRes.Errors[0].Detail, disCardRes.Errors[0].Code)
	}

	if disCardRes.Card.Enabled == true {
		t.Errorf("Card is not disabled")
	}
}
