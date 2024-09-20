package client

import (
	"time"

	"github.com/enuberepos/Rocket_SDK_Go/api"
)

func (c *defaultClient) GetIndicators(filters api.UsageFilters) (api.Indicator, error) {
	filters.Add("view", "resellers")
	return c.api.Indicators(filters)
}

func (c *defaultClient) GetIndicatorsPeriod(start time.Time, end time.Time) (api.Indicator, error) {
	filters := api.UsageFilters{
		Filters: api.KeyValueArray{
			{Key: "start_time", Value: start.Format(time.DateOnly)},
			{Key: "end_time", Value: end.Format(time.DateOnly)},
		},
	}

	return c.GetIndicators(filters)
}

func (c *defaultClient) GetIndicatorsMonth(y int, m time.Month) (api.Indicator, error) {
	start, end := getMonthPeriod(y, m)

	return c.GetIndicatorsPeriod(start, end)
}

func (c *defaultClient) GetIndicatorsCurrent() (api.Indicator, error) {
	return c.GetIndicatorsMonth(now.Year(), now.Month())
}
