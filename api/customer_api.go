package api

import (
	"fmt"

	"github.com/KoeInoue/square-go-sdk/models"
)

// CustomerApiInterface defines the interface for the CustomerApi
type CustomersApiInterface interface {
	CreateCustomer(models.CreateCustomerRequest) (*models.CreateCustomerResponse, error)
	DeleteCustomer(req models.DeleteCustomerRequest) (*models.DeleteCustomerResponse, error)
	RetrieveCustomer(cID string) (*models.RetrieveCustomerResponse, error)
}

// CustomerApi is a struct that implements CustomerApiInterface
type CustomersApi struct{}

// NewCustomerApi returns a new CustomerApi
func NewCustomerApi() *CustomersApi {
	return &CustomersApi{}
}

// CreateCustomer create customer
func (api *CustomersApi) CreateCustomer(req models.CreateCustomerRequest) (*models.CreateCustomerResponse, error) {
	res := models.CreateCustomerResponse{}
	return request(req, res, "/v2/customers", HTTP_POST)
}

// DeleteCustomer deletes a customer profile from a business.
func (api *CustomersApi) DeleteCustomer(req models.DeleteCustomerRequest) (*models.DeleteCustomerResponse, error) {
	res := models.DeleteCustomerResponse{}
	path := fmt.Sprintf("/v2/customers/%s?version=%d", req.CustomerId, req.Version)

	return request(req, res, path, HTTP_DELETE)
}

// RetrieveCustomer returns details for a single customer.
func (api *CustomersApi) RetrieveCustomer(cID string) (*models.RetrieveCustomerResponse, error) {
	res := models.RetrieveCustomerResponse{}
	path := fmt.Sprintf("/v2/customers/%s", cID)

	return request(struct{}{}, res, path, HTTP_GET)
}
