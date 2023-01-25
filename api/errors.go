package api

import (
	"errors"
	"strings"
)

type commonErrors struct {
	Errors *[]string `json:"errors"`
}

func (e *commonErrors) ToError() error {
	if e.Errors == nil {
		return nil
	}
	return errors.New(strings.Join(*e.Errors, ", "))
}
