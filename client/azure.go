package client

import (
	"net/url"

	"github.com/enuberepos/Rocket_SDK_Go/api"
)

type AzureClient struct {
	Client
	api *api.AzureAPI
}

func NewAzureClient(endpoint *url.URL) *AzureClient {
	api := api.NewAzureAPI(endpoint, api.Token{})
	return &AzureClient{
		Client: &defaultClient{api},
		api:    api,
	}
}
