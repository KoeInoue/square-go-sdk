package api

import (
	"github.com/KoeInoue/square-go-sdk/models"
)

// CustomerApiInterface defines the interface for the CustomerApi
type CustomersApiInterface interface {
	CreateCustomer(models.CreateCustomerRequest) (*models.CreateCustomerResponse, error)
}

// CustomerApi is a struct that implements CustomerApiInterface
type CustomersApi struct {}

// NewCustomerApi returns a new CustomerApi
func NewCustomerApi() *CustomersApi {
	return &CustomersApi{}
}

// CreateCustomer create customer
func (api *CustomersApi) CreateCustomer(req models.CreateCustomerRequest) (*models.CreateCustomerResponse, error) {
	res := models.CreateCustomerResponse{}
	return request(req, res, "/v2/customers", HTTP_POST)
}
