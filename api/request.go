package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/torotonnato/gobarebones/config"
)

func Request(method string, endPoint string, optData []byte, response interface{}) (int, error) {
	uri := config.Env.Region + endPoint
	optDataBuff := bytes.NewBuffer(optData)
	req, err := http.NewRequest(method, uri, optDataBuff)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("DD-API-KEY", config.Env.APIKey)
	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}
