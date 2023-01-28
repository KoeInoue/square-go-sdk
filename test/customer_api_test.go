package test

import (
	"reflect"
	"testing"

	"github.com/KoeInoue/square-go-sdk/models"
	"github.com/google/uuid"
)

func TestCreateAndDeleteCustomer(t *testing.T) {
	cases := []struct {
		Name        string
		Request     models.CreateCustomerRequest
		Expected    *models.CreateCustomerResponse
		IsSquareErr bool
		Error       []models.Error
	}{
		{
			Name: "Test with valid request values",
			Request: models.CreateCustomerRequest{
				EmailAddress:   "test@example.com",
				Nickname:       "Nickname",
				IdempotencyKey: uuid.New().String(),
				ReferenceId:    "Test",
			},
			Expected:    &models.CreateCustomerResponse{},
			IsSquareErr: false,
		},
		{
			Name: "Test with not enough request values",
			Request: models.CreateCustomerRequest{
				IdempotencyKey: uuid.New().String(),
			},
			Expected:    &models.CreateCustomerResponse{},
			IsSquareErr: true,
			Error: []models.Error{{
				Category: "INVALID_REQUEST_ERROR",
				Code:     "BAD_REQUEST",
				Detail:   "At least one of `given_name`, `family_name`, `company_name`, `email_address`, or `phone_number` is required for a customer.",
				Field:    "",
			}},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Helper()

			res, err := client.CustomerApi.CreateCustomer(test.Request)
			if err != nil {
				t.Errorf("got error, want nil: %s", err)
			}

			if reflect.ValueOf(res).Type() != reflect.ValueOf(test.Expected).Type() {
				t.Errorf("got %T, want %T", res, test.Expected)
			}

			if len(res.Errors) > 0 && !test.IsSquareErr {
				t.Errorf("Square Error: %#v, Code: %s", res.Errors[0].Detail, res.Errors[0].Code)
			}

			if test.IsSquareErr {
				if !reflect.DeepEqual(res.Errors, test.Error) {
					t.Errorf("got %#v, want %#v", res.Errors, test.Error)
				}
			}

			if len(res.Errors) == 0 && res.Customer.ID != "" {
				r, err := client.CustomerApi.DeleteCustomer(models.DeleteCustomerRequest{
					CustomerId: res.Customer.ID,
					Version:    res.Customer.Version,
				})
				if err != nil {
					t.Errorf("got error, want nil: %s", err)
				}
				if len(r.Errors) > 0 {
					t.Errorf("Square Error: %#v, Code: %s", r.Errors[0].Detail, r.Errors[0].Code)
				}
			}
		})
	}
}
