package core

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

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type KeyValueArray []KeyValuePair
