package api

import (
	"net/http"
)

type responseValidate struct {
	Valid *bool `json:"valid"`
	commonErrors
}

func Validate() (bool, error) {
	resp := responseValidate{}
	status, err := Request("GET", "/api/v1/validate", nil, &resp)
	if err != nil {
		return false, err
	}
	if status == http.StatusOK && resp.Valid != nil {
		return *resp.Valid, nil
	}
	if resp.Errors != nil {
		return false, resp.ToError()
	}
	return false, InvalidResponse{}
}
