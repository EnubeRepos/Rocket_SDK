package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/enuberepos/Rocket_SDK_Go/core"
)

var now = time.Now()

func getMonthPeriod(y int, m time.Month) (time.Time, time.Time) {
	f := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	l := time.Date(y, m+1, 1, 0, 0, 0, -1, time.UTC)
	return f, l
}

func (c *Client) GetUsage(filters core.UsageFilters) ([]core.Usage, error) {
	body, err := json.Marshal(filters)
	if err != nil {
		return []core.Usage{}, nil
	}

	c.endpoint.Path = "/api/v4/azure/usages"
	q := c.endpoint.Query()
	q.Add("timestamp", strconv.Itoa(int(now.Unix())))
	c.endpoint.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, c.endpoint.String(), bytes.NewReader(body))
	if err != nil {
		return []core.Usage{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token.Token))

	res, err := c.http.Do(req)
	if err != nil {
		return []core.Usage{}, err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return []core.Usage{}, err
	}

	if res.StatusCode != http.StatusOK {
		return []core.Usage{}, fmt.Errorf("API returned a non-200 code.\nCode %s\nBody:\n%s", res.Status, string(body))
	}

	var usage []core.Usage
	err = json.Unmarshal(body, &usage)

	return usage, err
}

func (c *Client) GetUsagePeriod(start time.Time, end time.Time) ([]core.Usage, error) {
	filters := core.UsageFilters{
		Filters: core.KeyValueArray{
			{Key: "start_time", Value: start.Format(time.DateOnly)},
			{Key: "end_time", Value: end.Format(time.DateOnly)},
			{Key: "view", Value: "resellers"},
		},
	}

	return c.GetUsage(filters)
}

func (c *Client) GetUsageMonth(y int, m time.Month) ([]core.Usage, error) {
	start, end := getMonthPeriod(y, m)

	return c.GetUsagePeriod(start, end)
}

func (c *Client) GetCurrentUsage() ([]core.Usage, error) {
	return c.GetUsageMonth(now.Year(), now.Month())
}
