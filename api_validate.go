package gobarebones

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type responseValidate struct {
	Valid *bool `json:"valid"`
	commonErrors
}

func parseValidate(resp *http.Response) (bool, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	respJSON := responseValidate{}
	err = json.Unmarshal(body, &respJSON) 
	if err != nil {
		return false, err
	}
	if resp.StatusCode == http.StatusOK && respJSON.Valid != nil {
		return *respJSON.Valid, nil
	}
	if respJSON.Errors != nil {
		return false, respJSON.ToError()
	}
	return false, errors.New("Invalid and unexpected response")
}

func apiValidate() (bool, error) {
	req, err := http.NewRequest("GET", config.Region + "/api/v1/validate", nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("DD-API-KEY", config.APIKey)
	resp, err := apiClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return parseValidate(resp)
}
