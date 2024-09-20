package client

import (
	"time"

	"github.com/enuberepos/Rocket_SDK_Go/api"
)

var now = time.Now()

func getMonthPeriod(y int, m time.Month) (time.Time, time.Time) {
	f := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	l := time.Date(y, m+1, 1, 0, 0, 0, -1, time.UTC)
	return f, l
}

func (c *defaultClient) GetUsage(filters api.UsageFilters) ([]api.Usage, error) {
	filters.Add("view", "resellers")
	return c.api.Usages(filters)
}

func (c *defaultClient) GetUsagePeriod(start time.Time, end time.Time) ([]api.Usage, error) {
	filters := api.UsageFilters{
		Filters: api.KeyValueArray{
			{Key: "start_time", Value: start.Format(time.DateOnly)},
			{Key: "end_time", Value: end.Format(time.DateOnly)},
		},
	}

	return c.GetUsage(filters)
}

func (c *defaultClient) GetUsageMonth(y int, m time.Month) ([]api.Usage, error) {
	start, end := getMonthPeriod(y, m)

	return c.GetUsagePeriod(start, end)
}

func (c *defaultClient) GetUsageCurrent() ([]api.Usage, error) {
	return c.GetUsageMonth(now.Year(), now.Month())
}
