package api

import (
	"net/url"
)

type AzureAPI struct {
	API
}

func NewAzureAPI(endpoint *url.URL, token Token) *AzureAPI {
	basepath, _ := url.JoinPath(endpoint.Path, "azure")
	authBasepath, _ := url.JoinPath(endpoint.Path, "auth")
	api := defaultApi{
		authBasepath: authBasepath,
		basepath:     basepath,
		endpoint:     endpoint,
		token:        token,
	}
	return &AzureAPI{&api}
}
