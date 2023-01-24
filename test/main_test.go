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

var testCustomerID = ""

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Failed to load .env file")
	}

	client = http.NewClient(square.Config[square.Sandbox]{
		AccessToken: os.Getenv("SQUARE_ACCESS_TOKEN"),
		Environment: square.Environments.Sandbox,
	})

	// Create a customer for testing
	resp, err := client.CustomerApi.CreateCustomer(models.CreateCustomerRequest{
		EmailAddress:   "test@example.com",
		Nickname:       "Nickname",
		IdempotencyKey: uuid.New().String(),
		ReferenceId:    "Test",
	})
	if err != nil {
		fmt.Println(err)
	}
	testCustomerID = resp.Customer.ID

	exitVal := m.Run()

	log.Println("Do stuff AFTER the tests!")
	// Delete the customer for testing
	client.CustomerApi.DeleteCustomer(models.DeleteCustomerRequest{
		CustomerId: testCustomerID,
	})

	os.Exit(exitVal)
}
