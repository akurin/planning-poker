package httpapi

import "net/url"

type Config interface {
	BasePath() *url.URL
}

type httpApiConfig struct {
	url *url.URL
}

func NewHttpApiConfig(url *url.URL) Config {
	return httpApiConfig{
		url: url,
	}
}

func (c httpApiConfig) BasePath() *url.URL {
	return c.url
}
