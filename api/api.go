package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type API interface {
	Login(user, passwd string) error
	Usages(UsageFilters) ([]Usage, error)
	Indicators(UsageFilters) (Indicator, error)
	Resellers(UsageFilters) ([]Reseller, error)
	CatalogTypes(UsageFilters) ([]CatalogType, error)
}

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type KeyValueArray []KeyValuePair
type Usage struct {
	ID           string         `json:"id"`
	ParentID     string         `json:"parent_id"`
	ParentName   string         `json:"parent_name,omitempty"`
	ResourceType string         `json:"resource_type"`
	Type         string         `json:"type"`
	Name         string         `json:"name"`
	Last         float64        `json:"last"`
	Actual       float64        `json:"actual"`
	Forecast     float64        `json:"forecast"`
	Leaf         bool           `json:"leaf"`
	Extra        map[string]any `json:"extra"`
	Data         map[string]any `json:"-"`
}
type UsageFilters struct {
	Filters KeyValueArray `json:"filters"`
}

func (f *UsageFilters) Add(key string, value string) {
	f.Filters = append(f.Filters, KeyValuePair{key, value})
}

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

type LabelValuePair struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type Reseller LabelValuePair
type CatalogType LabelValuePair

type Stack struct {
	Label           string                 `json:"label"`
	Value           float64                `json:"value"`
	Type            string                 `json:"type"`
	MarkupReference string                 `json:"markup_reference"`
	Extra           map[string]interface{} `json:"extra"`
}

type Tree struct {
	Name     string   `json:"name"`
	Value    *float64 `json:"value,omitempty"`
	Children []Tree   `json:"children,omitempty"`
}

func post[T any](filters UsageFilters, endpoint *url.URL, token string) (T, error) {
	var resValue T

	body, err := json.Marshal(filters)
	if err != nil {
		return resValue, err
	}

	req, err := http.NewRequest(http.MethodPost, endpoint.String(), bytes.NewReader(body))
	if err != nil {
		return resValue, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return resValue, err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resValue, err
	}

	if res.StatusCode != http.StatusOK {
		return resValue, fmt.Errorf("Non-200 code returned from api\ncode: %s\nbody:\n%s", res.Status, string(body))
	}

	err = json.Unmarshal(body, &resValue)
	return resValue, err
}
