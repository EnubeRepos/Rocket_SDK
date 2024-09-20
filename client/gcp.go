package client

import (
	"net/url"

	"github.com/enuberepos/Rocket_SDK_Go/api"
)

type GcpClient struct {
	Client
	api *api.GcpAPI
}

func NewGcpClient(endpoint *url.URL) *GcpClient {
	api := api.NewGcpAPI(endpoint, api.Token{})
	return &GcpClient{
		Client: &defaultClient{api},
		api:          api,
	}
}
