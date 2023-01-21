# Cards

```go
cardsApi := client.cardsApi
```

## Struct Name

`CardsApi`

## Methods

* [Create Card](../../doc/api/cards.md#create-card)

# Create Card

Adds a card on file to an existing merchant.

```ts
async createCard(
  body: CreateCardRequest,
  requestOptions?: RequestOptions
): Promise<ApiResponse<CreateCardResponse>>
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`CreateCardRequest`](../../doc/models/create-card-request.md) | Body, Required | An object containing the fields to POST for the request.<br><br>See the corresponding object definition for field details. |
| `requestOptions` | `RequestOptions \| undefined` | Optional | Pass additional request options. |

## Response Type

[`CreateCardResponse`](../../doc/models/create-card-response.md)

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

card := models.Card{
    CardholderName: "Amelia Earhart",
    BillingAddress: address,
    CustomerId: "VDKXEEKPJN48QDG3BGGFAK05P8",
    ReferenceId: "user-id-1",
}

body := models.CreateCardRequest{
  idempotencyKey: uuid.New().String(),
  sourceId: "cnon:uIbfJXhXETSP197M3GB",
  card: card,
};

res, err := client.CardApi.CreateCard(body)
```
