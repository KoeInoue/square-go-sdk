package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/KoeInoue/square-go-sdk"
	"github.com/KoeInoue/square-go-sdk/models"
)

const HTTP_POST = "POST"
const HTTP_GET = "GET"

type endpoint struct {
	path   string
	method string
}

// Api has config data for calling api
type Api struct {
	accessToken string
	envDomain   string
}

// ApiRequest is struct to call api
var ApiRequest = Api{
	accessToken: "",
	envDomain:   string(square.Environments.Sandbox),
}

// NewApi returns api client
func NewApi(accessToken string, envDomain string) {
	ApiRequest = Api{
		accessToken: accessToken,
		envDomain:   envDomain,
	}
}

// postRequest call http post request
func request[Req any, Res any](reqBody Req, res Res, path string, method string) (*Res, error) {
    var req *http.Request
	var err error
	if method == HTTP_POST {
		body, err := getByteBody(reqBody)
		if err != nil {
			return nil, err
		}
        req, err = http.NewRequest(method, ApiRequest.envDomain+path, body)

        if err != nil {
            return nil, err
        }
	} else if method == HTTP_GET {
        req, err = http.NewRequest(method, ApiRequest.envDomain+path, nil)
	}

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
		errRes := models.Error{}
		if err := json.Unmarshal(rBody, &errRes); err != nil {
			return nil, err
		}
		return nil, &errRes
	}

	return &res, nil
}

// getByteBody convert struct to *bytes.Reader
func getByteBody[T any](req T) (*bytes.Reader, error) {
	payloadBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(payloadBytes), nil
}

// setHeaders set header to *http.Request
func setHeaders(req *http.Request) {
	req.Header.Set("Square-Version", "2022-12-14")
	req.Header.Set("Authorization", "Bearer "+ApiRequest.accessToken)
	req.Header.Set("Content-Type", "application/json")
}
