package http

import (
	"github.com/KoeInoue/square-go-sdk"
	"github.com/KoeInoue/square-go-sdk/api"
)

// Client is http client
type Client struct {
    CustomerApi api.CustomerApiInterface
}

// NewClient returns Client struct pointer
func NewClient[T square.Env](config square.Config[T]) *Client {
    api.NewApi(config.AccessToken, string(config.Environment))

    return &Client{
        CustomerApi: api.NewCustomerApi(),
    }
}
