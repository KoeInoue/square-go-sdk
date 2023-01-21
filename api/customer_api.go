package api

import (
	"github.com/KoeInoue/square-go-sdk/models"
)

// CustomerApiInterface defines the interface for the CustomerApi
type CustomersApiInterface interface {
	CreateCustomer(models.CreateCustomerRequest) (*models.CreateCustomerResponse, error)
}

// CustomerApi is a struct that implements CustomerApiInterface
type CustomersApi struct {
	createCustomer endpoint
}

// NewCustomerApi returns a new CustomerApi
func NewCustomerApi() *CustomersApi {
	return &CustomersApi{
		createCustomer: endpoint{
			path:   "/v2/customers",
			method: "POST",
		},
	}
}

// CreateCustomer create customer
func (api *CustomersApi) CreateCustomer(req models.CreateCustomerRequest) (*models.CreateCustomerResponse, error) {
	res := models.CreateCustomerResponse{}
	return request(req, res, api.createCustomer.path, api.createCustomer.method)
}
