package api

import (
	"github.com/KoeInoue/square-go-sdk/models"
)

// CustomerApiInterface defines the interface for the CustomerApi
type CardApiInterface interface {
	CreateCard(models.CreateCardRequest) (*models.CreateCardResponse, error)
}

// CustomerApi is a struct that implements CustomerApiInterface
type CardsApi struct {
	createCard endpoint
}

// NewCustomerApi returns a new CustomerApi
func NewCardApi() *CardsApi {
	return &CardsApi{
		createCard: endpoint{
            path: "/v2/customers",
            method: "POST",
        },
	}
}

// CreateCard create card
func (api *CardsApi) CreateCard(req models.CreateCardRequest) (*models.CreateCardResponse, error) {
	res := models.CreateCardResponse{}
	return request(req, res, api.createCard.path, api.createCard.method)
}
