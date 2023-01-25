package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/KoeInoue/square-go-sdk"
	"github.com/KoeInoue/square-go-sdk/models"
)

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

// request call http post request
func request[Req any, Res any](reqBody Req, res Res, path string, method string) (*Res, error) {
	var req *http.Request
	var err error
	if method == HTTP_POST || method == HTTP_PUT || method == HTTP_PATCH {
		body, err := getByteBody(reqBody)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, ApiRequest.envDomain+path, body)

		if err != nil {
			return nil, err
		}
	} else if method == HTTP_GET || method == HTTP_DELETE {
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

func structToUrlQuery[T any](req T, u *url.URL) string {
	values := u.Query()
	rv := reflect.ValueOf(req)
	rt := rv.Type()
	// NumFieldでフィールド数を取得
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.FieldByName(field.Name) // value は interface です。
		kind := field.Type.Kind()

		tag := string(field.Tag.Get("json"))
		t := strings.Replace(tag, ",omitempty", "", -1)
		if kind == reflect.String && value.String() != "" {
			values.Add(t, value.String())
		}

		if kind == reflect.Bool {
			values.Add(t, strconv.FormatBool(value.Bool()))
		}
	}
	u.RawQuery = values.Encode()
	url := u.String()

	return url
}
