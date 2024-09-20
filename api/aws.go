package api

import (
	"net/url"
)

type AwsAPI struct {
	API
}

func NewAwsAPI(endpoint *url.URL, token Token) *AwsAPI {
	basepath, _ := url.JoinPath(endpoint.Path, "aws")
	authBasepath, _ := url.JoinPath(endpoint.Path, "auth")
	api := defaultApi{
		authBasepath: authBasepath,
		basepath:     basepath,
		endpoint:     endpoint,
		token:        token,
	}
	return &AwsAPI{&api}
}
