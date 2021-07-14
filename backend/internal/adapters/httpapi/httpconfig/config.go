package httpconfig

import "net/url"

type HttpApiConfig interface {
	BasePath() *url.URL
}

type httpApiConfig struct {
	url *url.URL
}

func NewHttpApiConfig(url *url.URL) HttpApiConfig {
	return httpApiConfig{
		url: url,
	}
}

func (c httpApiConfig) BasePath() *url.URL {
	return c.url
}
