package test

import (
	"log"
	"testing"

	"github.com/KoeInoue/square-go-sdk/models"
	"github.com/google/uuid"
)

func TestListCards(t *testing.T) {
	cardRes, err := client.CardApi.CreateCard(models.CreateCardRequest{
		SourceId:       "cnon:card-nonce-ok",
		IdempotencyKey: uuid.New().String(),
		Card: models.Card{
			CustomerId: testCustomerID,
		},
	})
	if err != nil {
		t.Error(err)
	}

	res, err := client.CardApi.ListCards(models.ListCardsRequest{
		CustomerId: cardRes.Card.CustomerId,
	})
	if err != nil {
		t.Error(err)
	}

	log.Printf("%#v", res)
}
