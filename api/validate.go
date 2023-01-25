package api

import (
	"net/http"
)

type responseValidate struct {
	Valid *bool `json:"valid"`
	APIErrors
}

func Validate() (bool, error) {
	resp := responseValidate{}
	status, err := Request("GET", "/api/v1/validate", nil, &resp)
	if err != nil {
		return false, err
	}
	if resp.HasError() {
		return false, resp.ToError()
	}
	if status == http.StatusOK && resp.Valid != nil {
		return *resp.Valid, nil
	}
	return false, InvalidResponse{}
}
