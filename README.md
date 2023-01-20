# Square Go SDK

<!-- # Badges -->

![go](https://img.shields.io/github/go-mod/go-version/KoeInoue/square-go-sdk)
![license](https://img.shields.io/github/license/KoeInoue/square-go-sdk)
![star](https://img.shields.io/github/stars/KoeInoue/square-go-sdk?style=social)

# Tags

`go` `square`

# Installation

```shell
go install github.com/KoeInoue/square-go-sdk
```

# Minimal Example

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

# Contributors

- [Koe Inoue](https://github.com/KoeInoue)

# Users

- [Koe Inoue](https://github.com/KoeInoue)

<!-- CREATED_BY_LEADYOU_README_GENERATOR -->
