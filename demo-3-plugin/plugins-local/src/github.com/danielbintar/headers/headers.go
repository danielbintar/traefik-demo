package headers

import (
	"context"
	"net/http"
)

type Config struct {
	CustomRequestHeaders  map[string]string `json:"customRequestHeaders", yaml:"customRequestHeaders"`
	CustomResponseHeaders map[string]string `json:"customRequestHeaders", yaml:"customRequestHeaders"`
}

func CreateConfig() *Config {
	return &Config{}
}

type HeadersPlugin struct {
	next                  http.Handler
	customRequestHeaders  map[string]string
	customResponseHeaders map[string]string
}

func New(_ context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	return &HeadersPlugin{
		next:                  next,
		customRequestHeaders:  config.CustomRequestHeaders,
		customResponseHeaders: config.CustomResponseHeaders,
	}, nil
}

func (s *HeadersPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for k, v := range s.customRequestHeaders {
		req.Header.Set(k, v)
	}

	for k, v := range s.customResponseHeaders {
		rw.Header().Set(k, v)
	}

	s.next.ServeHTTP(rw, req)
}
