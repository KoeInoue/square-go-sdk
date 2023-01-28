package api

import "github.com/KoeInoue/square-go-sdk/models"

// SubscriptionApiInterface defines the interface for the SubscriptionApi
type SubscriptionApiInterface interface {
	CreateSubscription(models.CreateSubscriptionRequest) (*models.CreateSubscriptionResponse, error)
	UpdateSubscription(string, models.UpdateSubscriptionRequest) (*models.UpdateSubscriptionResponse, error)
	CancelSubscription(string) (*models.CancelSubscriptionResponse, error)
}

// SubscriptionApi is a struct that implements SubscriptionApiInterface
type SubscriptionApi struct{}

// NewSubscriptionApi returns a new SubscriptionApi
func NewSubscriptionApi() *SubscriptionApi {
	return &SubscriptionApi{}
}

// CreateSubscription create new subscription
func (api *SubscriptionApi) CreateSubscription(req models.CreateSubscriptionRequest) (*models.CreateSubscriptionResponse, error) {
	res := models.CreateSubscriptionResponse{}
	return request(req, res, "/v2/subscriptions", HTTP_POST)
}

// UpdateSubscription update a subscription
func (api *SubscriptionApi) UpdateSubscription(subscriptionID string, req models.UpdateSubscriptionRequest) (*models.UpdateSubscriptionResponse, error) {
	res := models.UpdateSubscriptionResponse{}
	return request(req, res, "/v2/subscriptions/"+subscriptionID, HTTP_PUT)
}

// CancelSubscription schedules a CANCEL action to cancel an active subscription by setting the canceled_date field to
// the end of the active billing period and changing the subscription status from ACTIVE to CANCELED after this date.
func (api *SubscriptionApi) CancelSubscription(subscriptionID string) (*models.CancelSubscriptionResponse, error) {
	res := models.CancelSubscriptionResponse{}
	u := "/v2/subscriptions/" + subscriptionID + "/cancel"
	return request(struct{}{}, res, u, HTTP_POST)
}
