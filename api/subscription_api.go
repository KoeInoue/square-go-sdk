package api

import "github.com/KoeInoue/square-go-sdk/models"

// SubscriptionApiInterface defines the interface for the SubscriptionApi
type SubscriptionApiInterface interface {
	CreateSubscription(models.CreateSubscriptionRequest) (*models.CreateSubscriptionResponse, error)
}

// SubscriptionApi is a struct that implements SubscriptionApiInterface
type SubscriptionApi struct {
	createSubscription endpoint
}

// NewSubscriptionApi returns a new SubscriptionApi
func NewSubscriptionApi() *SubscriptionApi {
	return &SubscriptionApi{
		createSubscription: endpoint{
			path:   "/v2/subscriptions",
			method: "POST",
		},
	}
}

// CreateSubscription create new subscription
func (api *SubscriptionApi) CreateSubscription(req models.CreateSubscriptionRequest) (*models.CreateSubscriptionResponse, error) {
	res := models.CreateSubscriptionResponse{}
	return request(req, res, api.createSubscription.path, api.createSubscription.method)
}
