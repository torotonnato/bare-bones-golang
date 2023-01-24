package gobarebones

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type responseSeries struct {
	commonErrors
}

func parseSeries(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respJSON := responseSeries{}
	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		fmt.Println("error json:", respJSON)
		return err
	}
	fmt.Println("body", string(body))
	if resp.StatusCode == http.StatusAccepted && len(*respJSON.Errors) == 0 {
		return nil
	}
	if respJSON.Errors != nil {
		return respJSON.ToError()
	}
	return errors.New("invalid and unexpected response")
}

type series struct {
	Series []metric `json:"series"`
}

func ApiSeries(m *metric) error {
	s := series{}
	s.Series = append(s.Series, *m)
	fmt.Println(s)

	mData, err := json.Marshal(&s)
	fmt.Println("asdasdasd", err, string(mData))
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", config.Region+"/api/v2/series", bytes.NewBuffer(mData))
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DD-API-KEY", config.APIKey)
	resp, err := apiClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return parseSeries(resp)
}
