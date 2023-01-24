package models

import (
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type ListCardsRequest struct {
	Cursor          string `json:"cursor,omitempty"`
	CustomerId      string `json:"customer_id,omitempty"`
	IncludeDisabled bool   `json:"include_disabled,omitempty"`
	ReferenceId     string `json:"reference_id,omitempty"`
	SortOrder       string `json:"sort_order,omitempty"`
}

func (m *ListCardsRequest) GetURLWithQuery(u *url.URL) string {
	values := u.Query()
	rv := reflect.ValueOf(*m)
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
	log.Printf("%s", url)

	return url
}
