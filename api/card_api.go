package api

import (
	"github.com/KoeInoue/square-go-sdk/models"
)

// CardApiInterface defines the interface for the CustomerApi
type CardApiInterface interface {
	CreateCard(models.CreateCardRequest) (*models.CreateCardResponse, error)
}

// CardsApi is a struct that implements CustomerApiInterface
type CardsApi struct {
	createCard endpoint
}

// NewCardApi returns a new CustomerApi
func NewCardApi() *CardsApi {
	return &CardsApi{
		createCard: endpoint{
			path:   "/v2/cards",
			method: "POST",
		},
	}
}

// CreateCard create new card
func (api *CardsApi) CreateCard(req models.CreateCardRequest) (*models.CreateCardResponse, error) {
	res := models.CreateCardResponse{}
	return request(req, res, api.createCard.path, api.createCard.method)
}
