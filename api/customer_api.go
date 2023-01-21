package api

import (
	"github.com/KoeInoue/square-go-sdk/models"
)

// CustomerApiInterface defines the interface for the CustomerApi
type CustomerApiInterface interface {
	CreateCustomer(models.CreateCustomerRequest) (*models.CreateCustomerResponse, error)
}

// CustomerApi is a struct that implements CustomerApiInterface
type CustomerApi struct {
	createCustomerPath string
}

// NewCustomerApi returns a new CustomerApi
func NewCustomerApi() *CustomerApi {
	return &CustomerApi{
		createCustomerPath: "/v2/customers",
	}
}

// CreateCustomer create customer
func (api *CustomerApi) CreateCustomer(req models.CreateCustomerRequest) (*models.CreateCustomerResponse, error) {
	res := models.CreateCustomerResponse{}
	return postRequest(req, res, api.createCustomerPath)
}
