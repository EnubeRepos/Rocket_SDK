package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

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

func (a *defaultApi) Usages(filters UsageFilters) ([]Usage, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "usages")

	body, err := json.Marshal(filters)
	if err != nil {
		return []Usage{}, err
	}

	req, err := http.NewRequest(http.MethodPost, a.endpoint.String(), bytes.NewReader(body))
	if err != nil {
		return []Usage{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token.Token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Usage{}, err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return []Usage{}, err
	}

	if res.StatusCode != http.StatusOK {
		return []Usage{}, err
	}

	var usage []Usage
	err = json.Unmarshal(body, &usage)

	return usage, err
}
