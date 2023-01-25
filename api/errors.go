package api

import (
	"strings"
)

type GenericError struct {
	msg string
}

func (e GenericError) Error() string {
	return e.msg
}

type BadRequest struct{}

func (e BadRequest) Error() string {
	return "bad request"
}

type InvalidResponse struct{}

func (e InvalidResponse) Error() string {
	return "invalid response from DataDog server"
}

// Common API error oject definition
type APIErrors struct {
	Errors *[]string `json:"errors"` //optional
}

func (e *APIErrors) HasError() bool {
	return e.Errors != nil
}

func (e *APIErrors) ToError() error {
	if e.Errors == nil {
		return nil
	}
	err := *e.Errors //Alias
	if len(err) == 1 {
		if err[0] == "Bad request" {
			return BadRequest{}
		}
	}
	return GenericError{
		strings.Join(*e.Errors, ", "),
	}
}
