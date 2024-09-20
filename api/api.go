package api

type API interface {
	Login(user, passwd string) error
	Usages(filters UsageFilters) ([]Usage, error)
	Indicators(filters UsageFilters) (Indicator, error)
}
