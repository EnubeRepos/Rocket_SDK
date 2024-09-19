package rocketsdkgo

import (
	"net/url"

	"github.com/enuberepos/Rocket_SDK_Go/client"
)

func Login(endpoint *url.URL, user string, passwd string) (*client.Client, error) {
	c := client.NewClient(endpoint)

	err := c.Login(user, passwd)
	if err != nil {
		return nil, err
	}

	return c, nil
}
