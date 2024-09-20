package client

import (
	"net/url"

	"github.com/enuberepos/Rocket_SDK_Go/api"
)

type AwsClient struct {
	Client
	api *api.AwsAPI
}

func NewAwsClient(endpoint *url.URL) *AwsClient {
	api := api.NewAwsAPI(endpoint, api.Token{})
	return &AwsClient{
		Client: &defaultClient{api},
		api:          api,
	}
}
