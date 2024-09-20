package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Indicator struct {
	Type             string   `json:"type,omitempty"`
	Last             float64  `json:"last"`
	LastExchangeRate *float64 `json:"lastExchangeRate"`
	Actual           float64  `json:"actual"`
	ExchangeRate     *float64 `json:"exchangeRate"`
	Forecast         float64  `json:"forecast"`
	LastQuantity     float64  `json:"last_quantity"`
	ActualQuantity   float64  `json:"actual_quantity"`
}

func (a *defaultApi) Indicators(filters UsageFilters) (Indicator, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "indicators")

	body, err := json.Marshal(filters)
	if err != nil {
		return Indicator{}, err
	}

	req, err := http.NewRequest(http.MethodPost, a.endpoint.String(), bytes.NewReader(body))
	if err != nil {
		return Indicator{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token.Token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Indicator{}, err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return Indicator{}, err
	}

	if res.StatusCode != http.StatusOK {
		return Indicator{}, err
	}

	var indicator Indicator
	err = json.Unmarshal(body, &indicator)

	return indicator, err
}
