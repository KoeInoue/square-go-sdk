package test

import (
	"os"
	"testing"
	"time"

	"github.com/KoeInoue/square-go-sdk/models"
	"github.com/google/uuid"
)

func TestSubscriptionApi(t *testing.T) {
	cRes, err := client.CatalogApi.ListCatalog(models.ListCatalogRequest{
		Types: "SUBSCRIPTION_PLAN",
	})
	if err != nil {
		t.Error(err)
	}
	if len(cRes.Errors) > 0 {
		t.Errorf("Square Error: %#v, Code: %s", cRes.Errors[0].Detail, cRes.Errors[0].Code)
	}

	csRes, err := client.SubscriptionApi.CreateSubscription(models.CreateSubscriptionRequest{
		IdempotencyKey: uuid.New().String(),
		CustomerId:     testCustomer.ID,
		PlanId:         cRes.Objects[0].ID,
		LocationId:     os.Getenv("SQUARE_LOCATION_ID"),
		PriceOverrideMoney: models.Money{
			Amount:   1000,
			Currency: "JPY",
		},
		StartDate: time.Now().Format("2006-01-02"),
		CardId:    testCustomer.Cards[0].ID,
		Timezone:  "Japan",
		Source: models.SubscriptionSource{
			Name: "API",
		},
	})
	if err != nil {
		t.Error(err)
	}
	if len(csRes.Errors) > 0 {
		t.Errorf("Square Error: %#v, Code: %s Field: %s", csRes.Errors[0].Detail, csRes.Errors[0].Code, csRes.Errors[0].Field)
	}

	usRes, err := client.SubscriptionApi.UpdateSubscription(csRes.Subscription.ID, models.UpdateSubscriptionRequest{
		Subscription: models.Subscription{
			PriceOverrideMoney: models.Money{
				Amount:   10000,
				Currency: "JPY",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	if len(usRes.Errors) > 0 {
		t.Errorf("Square Error: %#v, Code: %s", usRes.Errors[0].Detail, usRes.Errors[0].Code)
	}

	canSRes, err := client.SubscriptionApi.CancelSubscription(csRes.Subscription.ID)
	if err != nil {
		t.Error(err)
	}
	if len(canSRes.Errors) > 0 {
		t.Errorf("Square Error: %#v, Code: %s", canSRes.Errors[0].Detail, canSRes.Errors[0].Code)
	}
}
