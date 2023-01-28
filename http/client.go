package http

import (
	"log"

	"github.com/KoeInoue/square-go-sdk"
	"github.com/KoeInoue/square-go-sdk/api"
)

// Client is http client
type Client struct {
	CustomerApi     api.CustomersApiInterface
	CardApi         api.CardApiInterface
	CatalogApi      api.CatalogApiInterface
	SubscriptionApi api.SubscriptionApiInterface
}

// NewClient returns Client struct pointer
func NewClient[T square.Env](config square.Config[T]) *Client {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	api.NewApi(config.AccessToken, string(config.Environment))

	return &Client{
		CustomerApi:     api.NewCustomerApi(),
		CardApi:         api.NewCardApi(),
		CatalogApi:      api.NewCatalogApi(),
		SubscriptionApi: api.NewSubscriptionApi(),
	}
}
