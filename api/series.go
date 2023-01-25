package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/torotonnato/gobarebones/model"
)

type responseSeries struct {
	commonErrors
}

func Series(s *model.Series) error {
	payload, err := json.Marshal(&s)
	if err != nil {
		return err
	}
	resp := responseSeries{}
	status, err := Request("POST", "/api/v2/series", payload, &resp)
	if err != nil {
		return err
	}
	if (status == http.StatusAccepted) && (resp.Errors == nil) {
		return nil
	}
	if resp.Errors != nil {
		return resp.ToError()
	}
	return errors.New("invalid and unexpected response")
}
