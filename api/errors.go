package api

import (
	"strings"
)

const (
	Generic = iota + 1
	BadRequest
	InvalidResponse
)

type Error struct {
	Code   int
	AuxMsg string
}

func (e Error) Error() string {
	switch e.Code {
	case Generic:
		return e.AuxMsg
	case BadRequest:
		return "bad request"
	case InvalidResponse:
		return "invalid response from DataDog server"
	}
	return "unknown error"
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
			return Error{Code: BadRequest}
		}
	}
	return Error{
		Code:   Generic,
		AuxMsg: strings.Join(*e.Errors, ", "),
	}
}
