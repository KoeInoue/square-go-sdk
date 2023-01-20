package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/KoeInoue/square-go-sdk"
)

type Api struct {
    accessToken string
    envDomain string
}

var ApiRequest = Api{
    accessToken: "",
    envDomain: string(square.Environments.Sandbox),
}


func NewApi(accessToken string, envDomain string) {
    ApiRequest = Api{
        accessToken: accessToken,
        envDomain: envDomain,
    }
}

type SquareApiErrRes struct {
    Errors []struct{
        Category string `json:"category"`
        Code string `json:"code"`
        Detail string `json:"detail"`
    } `json:"errors"`
}
func (e *SquareApiErrRes) Error() string {
    return fmt.Sprintf("Square API Error: \x1b[31m%#v\x1b[0m", e)
}

func postRequest[Req comparable, Res comparable](reqBody Req, res Res, path string) (*Res, error){
    body, err := getByteBody(reqBody)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("POST", ApiRequest.envDomain + path, body)
    if err != nil {
        return nil, err
    }
    setHeaders(req)

    resp, err := http.DefaultClient.Do(req)

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

	rBody, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal(rBody, &res); err != nil {
        errRes := SquareApiErrRes{}
        if err := json.Unmarshal(rBody, &errRes); err != nil {
            return nil, err
        }
        return nil, &errRes
	}

    return &res, nil
}

func getByteBody[T comparable](req T) (*bytes.Reader, error) {
    payloadBytes, err := json.Marshal(req)
    if err != nil {
        return nil, err
    }

    return bytes.NewReader(payloadBytes), nil
}

func setHeaders(req *http.Request) {
    req.Header.Set("Square-Version", "2022-12-14")
    req.Header.Set("Authorization", "Bearer " + ApiRequest.accessToken)
    req.Header.Set("Content-Type", "application/json")
}
