# Customers

```go
customersApi := client.customersApi;
```

## Struct Name

`CustomersApi`

## Methods

* [Create Customer](../../doc/api/customers.md#create-customer)

# Create Customer

Creates a new customer for a business.

You must provide at least one of the following values in your request to this
endpoint:

- `given_name`
- `family_name`
- `company_name`
- `email_address`
- `phone_number`

```go
CreateCustomer(models.CreateCustomerRequest) (*models.CreateCustomerResponse, error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`CreateCustomerRequest`](../../doc/models/create-customer-request.md) | Body, Required | An struct containing the fields to POST for the request.<br><br>See the corresponding object definition for field details. |

## Response Type

[`CreateCustomerResponse`](../../doc/models/create-customer-response.md)

## Example Usage

```go
address := models.Address{
    AddressLine1: "500 Electric Ave",
    addressLine2 = "Suite 600",
    locality = "New York",
    administrativeDistrictLevel1 = "NY",
    postalCode = "10003",
    country = "US",
}

body := models.CreateCustomerRequest{
    EmailAddress: "test@example.com",
    Nickname:    "your nickname",
    Address: address,
    IdempotencyKey: uuid.New().String(),
    Note: "a customer",
}

res, err := client.CustomerApi.CreateCustomer(body)
```
