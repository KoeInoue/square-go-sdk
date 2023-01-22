package api

import (
	"github.com/KoeInoue/square-go-sdk/models"
)

// CardApiInterface defines the interface for the CustomerApi
type CardApiInterface interface {
	CreateCard(models.CreateCardRequest) (*models.CreateCardResponse, error)
}

// CardsApi is a struct that implements CustomerApiInterface
type CardsApi struct {}

// NewCardApi returns a new CustomerApi
func NewCardApi() *CardsApi {
	return &CardsApi{}
}

// CreateCard create new card
func (api *CardsApi) CreateCard(req models.CreateCardRequest) (*models.CreateCardResponse, error) {
	res := models.CreateCardResponse{}
	return request(req, res, "/v2/cards", HTTP_POST)
}
