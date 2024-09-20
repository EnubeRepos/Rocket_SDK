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

}

type defaultClient struct {
	api api.API
}

func (c *defaultClient) Login(user string, passwd string) error {
	return c.api.Login(user, passwd)
}
