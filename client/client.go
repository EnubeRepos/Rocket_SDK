package client

import (
	"time"

	"github.com/enuberepos/Rocket_SDK_Go/api"
)

type Client interface {
	Login(user, passwd string) error

	GetUsage(filters api.UsageFilters) ([]api.Usage, error)
	GetUsagePeriod(start, end time.Time) ([]api.Usage, error)
	GetUsageMonth(year int, month time.Month) ([]api.Usage, error)
	GetUsageCurrent() ([]api.Usage, error)

	GetIndicators(filters api.UsageFilters) (api.Indicator, error)
	GetIndicatorsPeriod(start, end time.Time) (api.Indicator, error)
	GetIndicatorsMonth(year int, month time.Month) (api.Indicator, error)
	GetIndicatorsCurrent() (api.Indicator, error)
}

type defaultClient struct {
	api api.API
}

func (c *defaultClient) Login(user string, passwd string) error {
	return c.api.Login(user, passwd)
}
