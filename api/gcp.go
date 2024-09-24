package api

import (
	"net/url"
)

type GcpAPI struct {
	API
}

func NewGcpAPI(endpoint *url.URL, token Token) *GcpAPI {
	e := *endpoint
	endpoint = &e
	basepath, _ := url.JoinPath(endpoint.Path, "gcp")
	authBasepath, _ := url.JoinPath(endpoint.Path, "auth")
	api := defaultApi{
		authBasepath: authBasepath,
		basepath:     basepath,
		endpoint:     endpoint,
		token:        token,
	}
	return &GcpAPI{&api}
}
