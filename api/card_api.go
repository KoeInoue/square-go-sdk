package api

import (
	"net/url"

	"github.com/KoeInoue/square-go-sdk/models"
)

// CardApiInterface defines the interface for the CustomerApi
type CardApiInterface interface {
	CreateCard(models.CreateCardRequest) (*models.CreateCardResponse, error)
	ListCards(req models.ListCardsRequest) (*models.ListCardsResponse, error)
	RetrieveCard(cardID string) (*models.RetrieveCardResponse, error)
	DisableCard(cardID string) (*models.DisableCardResponse, error)
}

// CardsApi is a struct that implements CustomerApiInterface
type CardsApi struct{}

// NewCardApi returns a new CustomerApi
func NewCardApi() *CardsApi {
	return &CardsApi{}
}

// CreateCard create new card
func (api *CardsApi) CreateCard(req models.CreateCardRequest) (*models.CreateCardResponse, error) {
	res := models.CreateCardResponse{}
	return request(req, res, "/v2/cards", HTTP_POST)
}

func (api *CardsApi) ListCards(req models.ListCardsRequest) (*models.ListCardsResponse, error) {
	res := models.ListCardsResponse{}

	u, err := url.Parse("/v2/cards")
	if err != nil {
		return nil, err
	}

	url := structToUrlQuery(req, u)

	return request(req, res, url, HTTP_GET)
}

func (api *CardsApi) RetrieveCard(cardID string) (*models.RetrieveCardResponse, error) {
	res := models.RetrieveCardResponse{}
	return request(struct{}{}, res, "/v2/cards/"+cardID, HTTP_GET)
}

func (api *CardsApi) DisableCard(cardID string) (*models.DisableCardResponse, error) {
	res := models.DisableCardResponse{}
	u := "/v2/cards/" + cardID + "/disable"
	return request(struct{}{}, res, u, HTTP_POST)
}
