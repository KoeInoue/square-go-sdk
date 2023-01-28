package test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/KoeInoue/square-go-sdk"
	"github.com/KoeInoue/square-go-sdk/http"
	"github.com/KoeInoue/square-go-sdk/models"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var client *http.Client

var testCustomer = models.Customer{}

func TestMain(m *testing.M) {
	log.Println("-------- Do stuff BEFORE the tests!")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Failed to load .env file")
	}

	client = http.NewClient(square.Config[square.Sandbox]{
		AccessToken: os.Getenv("SQUARE_ACCESS_TOKEN"),
		Environment: square.Environments.Sandbox,
	})

	resp, err := createTestCustomer()

	log.Println("Test customer created: ", resp.Customer.ID)
	if err != nil {
		fmt.Println(err)
	}

	cRes, err := createTestCustomerCard(resp.Customer.ID)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Test customer card created: ", cRes.Card.ID)

	newC, err := retrieveTestCustomer(resp.Customer.ID)
	if err != nil {
		fmt.Println(err)
	}
	testCustomer = newC.Customer

	log.Println("-------- Start the tests!")
	exitVal := m.Run()

	log.Println("-------- Do stuff AFTER the tests!")

	// Delete the customer for testing
	deleteTestCustomer(testCustomer.ID)
	log.Println("Test customer deleted")

	os.Exit(exitVal)
}

func createTestCustomer() (*models.CreateCustomerResponse, error) {
	// Create a customer for testing
	return client.CustomerApi.CreateCustomer(models.CreateCustomerRequest{
		EmailAddress:   "test@example.com",
		Nickname:       "Nickname",
		IdempotencyKey: uuid.New().String(),
		ReferenceId:    "Test",
	})
}

func createTestCustomerCard(cID string) (*models.CreateCardResponse, error) {
	// Create a customer for testing
	return client.CardApi.CreateCard(models.CreateCardRequest{
		IdempotencyKey: uuid.New().String(),
		Card: models.Card{
			CustomerId: cID,
		},
		SourceId: "cnon:card-nonce-ok",
	})
}

func retrieveTestCustomer(cID string) (*models.RetrieveCustomerResponse, error) {
	return client.CustomerApi.RetrieveCustomer(cID)
}

func deleteTestCustomer(cID string) {
	client.CustomerApi.DeleteCustomer(models.DeleteCustomerRequest{
		CustomerId: cID,
	})
}
