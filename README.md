# Square Go SDK

<!-- # Badges -->

![go](https://img.shields.io/github/go-mod/go-version/KoeInoue/square-go-sdk)
![license](https://img.shields.io/github/license/KoeInoue/square-go-sdk)
![star](https://img.shields.io/github/stars/KoeInoue/square-go-sdk?style=social)
[![Go Test](https://github.com/KoeInoue/square-go-sdk/actions/workflows/go-test.yml/badge.svg)](https://github.com/KoeInoue/square-go-sdk/actions/workflows/go-test.yml)

## Tags

`go` `square`

## Installation

```shell
go install github.com/KoeInoue/square-go-sdk
```

or

import the package in a file where you want to use the package
```go
import "github.com/KoeInoue/square-go-sdk"
```

then run
```go
go mod tidy
```

## Minimal Example

1. Set environment variables
```shell
export SQUARE_ACCESS_TOKEN=yourSandboxAccessToken
export ENV=dev
```
※ ENV is a variable that just identifies the environment

2. Write code

```go
package main

import (
    "os"

	"github.com/KoeInoue/square-go-sdk"
	"github.com/KoeInoue/square-go-sdk/http"
	"github.com/KoeInoue/square-go-sdk/models"
)

func getClient() *http.Client {
    envType := os.Getenv("ENV")
    accessToken := os.Getenv("SQUARE_ACCESS_TOKEN")

    if envType != "production" {
        return http.NewClient[square.Sandbox](square.Config[square.Sandbox]{
            AccessToken: accessToken,
            Environment: square.Environments.Sandbox,
        })
    } else {
        return http.NewClient[square.Production](square.Config[square.Production]{
            AccessToken: accessToken,
            Environment: square.Environments.Production,
        })
    }
}

func main() {
    client := getClient()
    client.CustomerApi.CreateCustomer()
}
```

## SDK Reference

### Customers
* [Customers]

## Contributors

- [Koe Inoue](https://github.com/KoeInoue)


<!-- CREATED_BY_LEADYOU_README_GENERATOR -->
