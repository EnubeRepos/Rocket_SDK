package rocketsdkgo

import (
	"net/url"

	"github.com/enuberepos/Rocket_SDK_Go/client"
)

func LoginAws(endpoint *url.URL, user, passwd string) (*client.AwsClient, error) {
	c := client.NewAwsClient(endpoint)

	err := c.Login(user, passwd)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func LoginAzure(endpoint *url.URL, user, passwd string) (*client.AzureClient, error) {
	c := client.NewAzureClient(endpoint)

	err := c.Login(user, passwd)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func LoginGcp(endpoint *url.URL, user, passwd string) (*client.GcpClient, error) {
	c := client.NewGcpClient(endpoint)

	err := c.Login(user, passwd)
	if err != nil {
		return nil, err
	}

	return c, nil
}
