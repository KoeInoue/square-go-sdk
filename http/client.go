package http

import (
	"github.com/KoeInoue/square-go-sdk"
	"github.com/KoeInoue/square-go-sdk/api"
)

// Client is http client
type Client struct {
	config      square.Config
	customerApi api.CustomerApiInterface
}

// NewClient returns Client struct pointer
func NewClient(config square.Config) *Client {
	api.NewApi(config.AccessToken, string(config.Environment.Sandbox))
	return &Client{
		config:      config,
		customerApi: api.NewCustomerApi(),
	}
}
