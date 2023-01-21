package models

import "fmt"

// Error represents an error encountered during a request to the Connect API.
type Error struct {
	Category string `json:"category"`
	Code     string `json:"code"`
	Detail   string `json:"detail"`
	Field    string `json:"field"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Square API Error: \x1b[31m%#v\x1b[0m", e)
}
