package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	token    token
	endpoint *url.URL
	http     *http.Client
}

type token struct {
	ExpiresAt        int       `json:"expiresAt"`
	RefreshExpiresAt time.Time `json:"refreshExpiresAt"`
	RefreshToken     string    `json:"refreshToken"`
	Token            string    `json:"token"`
}

func NewClient(endpoint *url.URL) *Client {
	return &Client{
		endpoint: endpoint,
		http:     http.DefaultClient,
	}
}

func (c *Client) Login(user string, passwd string) error {
	input := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{user, passwd}

	i, _ := json.Marshal(input)

	c.endpoint.Path = "/api/v4/auth/login"
	res, err := c.http.Post(c.endpoint.String(), "application/json", bytes.NewReader(i))
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var token token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return err
	}

	c.token = token
	return nil
}
