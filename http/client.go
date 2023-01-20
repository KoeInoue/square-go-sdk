package http

import (
	"github.com/KoeInoue/square-go-sdk"
	"github.com/KoeInoue/square-go-sdk/api"
)

type Client struct {
    CustomerApi api.CustomerApiInterface
}

func NewClient[T square.Env](config square.Config[T]) *Client {
    api.NewApi(config.AccessToken, string(config.Environment))

    return &Client{
        CustomerApi: api.NewCustomerApi(),
    }
}
