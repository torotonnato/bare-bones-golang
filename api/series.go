package api

import (
	"encoding/json"
	"net/http"

	"github.com/torotonnato/gobarebones/model"
)

type responseSeries struct {
	APIErrors
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
	if resp.HasError() {
		return resp.ToError()
	}
	if status == http.StatusAccepted {
		return nil
	}
	return Error{Code: InvalidResponse}
}
