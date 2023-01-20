package api

import (
	"github.com/KoeInoue/square-go-sdk/models"
)

type CustomerApiInterface interface {
    CreateCustomer(models.CreateCustomerRequest) (*models.CreateCustomerResponse, error)
}

type CustomerApi struct {
    createCustomerPath string
}

func NewCustomerApi() * CustomerApi {
    return &CustomerApi{
        createCustomerPath: "/v2/customers",
    }
}

func (api *CustomerApi) CreateCustomer(req models.CreateCustomerRequest) (*models.CreateCustomerResponse, error) {
    res := models.CreateCustomerResponse{}
    return postRequest(req, res, api.createCustomerPath)
}
