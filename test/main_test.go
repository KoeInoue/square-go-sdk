package test

import (
	"log"
	"os"
	"testing"

	"github.com/KoeInoue/square-go-sdk"
	"github.com/KoeInoue/square-go-sdk/http"
	"github.com/joho/godotenv"
)

var client *http.Client

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

	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}
