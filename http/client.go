package http

import (
	"github.com/KoeInoue/square-go-sdk"
	"github.com/KoeInoue/square-go-sdk/api"
)

type Client struct {
    CustomerApi api.CustomerApiInterface
}

func NewClient(config square.Config) *Client {
    api.NewApi(config.AccessToken, string(config.Environment.Sandbox))
    return &Client{
        CustomerApi: api.NewCustomerApi(),
    }
}
