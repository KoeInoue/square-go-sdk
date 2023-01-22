package api

import "github.com/KoeInoue/square-go-sdk/models"

// SubscriptionApiInterface defines the interface for the SubscriptionApi
type SubscriptionApiInterface interface {
	CreateSubscription(models.CreateSubscriptionRequest) (*models.CreateSubscriptionResponse, error)
}

// SubscriptionApi is a struct that implements SubscriptionApiInterface
type SubscriptionApi struct {}

// NewSubscriptionApi returns a new SubscriptionApi
func NewSubscriptionApi() *SubscriptionApi {
	return &SubscriptionApi{}
}

// CreateSubscription create new subscription
func (api *SubscriptionApi) CreateSubscription(req models.CreateSubscriptionRequest) (*models.CreateSubscriptionResponse, error) {
	res := models.CreateSubscriptionResponse{}
	return request(req, res, "/v2/subscriptions", HTTP_POST)
}
