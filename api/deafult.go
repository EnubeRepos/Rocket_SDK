package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Token struct {
	ExpiresAt        int       `json:"expiresAt"`
	RefreshExpiresAt time.Time `json:"refreshExpiresAt"`
	RefreshToken     string    `json:"refreshToken"`
	Token            string    `json:"token"`
}

type defaultApi struct {
	endpoint     *url.URL
	basepath     string
	authBasepath string
	token        Token
}

func (a *defaultApi) Login(user string, passwd string) error {
	var err error

	a.endpoint.Path, err = url.JoinPath(a.authBasepath, "login")
	if err != nil {
		return err
	}

	input := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{user, passwd}

	i, _ := json.Marshal(input)

	res, err := http.Post(a.endpoint.String(), "application/json", bytes.NewReader(i))
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	} else if res.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned a non-200 code\nCode %s\nBody\n%s", res.Status, string(body))
	}

	var token Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return err
	}

	a.token = token
	return nil
}

func (a *defaultApi) Usages(f UsageFilters) ([]Usage, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "usages")
	return post[[]Usage](f, a.endpoint, a.token.Token)
}

func (a *defaultApi) Indicators(f UsageFilters) (Indicator, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "indicators")
	return post[Indicator](f, a.endpoint, a.token.Token)
}

func (a *defaultApi) Resellers(f UsageFilters) ([]Reseller, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "resellers")
	return post[[]Reseller](f, a.endpoint, a.token.Token)
}

func (a *defaultApi) CatalogTypes(f UsageFilters) ([]CatalogType, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "resellers")
	return post[[]CatalogType](f, a.endpoint, a.token.Token)
}

func (a *defaultApi) ByCategory(f UsageFilters) (Tree, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "by-category")
	return post[Tree](f, a.endpoint, a.token.Token)
}
func (a *defaultApi) ByDay(f UsageFilters) ([]Stack, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "by-day")
	return post[[]Stack](f, a.endpoint, a.token.Token)
}

func (a *defaultApi) ByMonth(f UsageFilters) ([]Stack, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "by-month")
	return post[[]Stack](f, a.endpoint, a.token.Token)
}

func (a *defaultApi) ByPublisher(f UsageFilters) ([]Stack, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "by-publisher")
	return post[[]Stack](f, a.endpoint, a.token.Token)
}

func (a *defaultApi) ByResource(f UsageFilters) ([]Usage, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "by-publisher")
	return post[[]Usage](f, a.endpoint, a.token.Token)
}

func (a *defaultApi) ByTag(f UsageFilters) (Tree, error) {
	a.endpoint.Path, _ = url.JoinPath(a.basepath, "by-tag")
	return post[Tree](f, a.endpoint, a.token.Token)
}
