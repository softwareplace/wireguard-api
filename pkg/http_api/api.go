package http_api

import "fmt"

type ApiConfig struct {
	Host               string
	Path               string
	Headers            map[string]string
	Query              map[string]string
	Body               any
	ExpectedStatusCode int
}

type Api[T any] interface {
	Get(config *ApiConfig) (*T, error)
	Post(config *ApiConfig) (*T, error)
	Put(config *ApiConfig) (*T, error)
	Delete(config *ApiConfig) (*T, error)
	Patch(config *ApiConfig) (*T, error)
	Head(config *ApiConfig) (*T, error)
}

func NewApi[T any](response T) Api[T] {
	i := new(_impl[T])
	i.response = response
	return i
}

type _impl[T any] struct {
	response T
}

func Config(host string) *ApiConfig {
	config := &ApiConfig{}
	config.Host = host
	config.Path = ""
	config.Headers = map[string]string{}
	config.Query = map[string]string{}
	config.Body = nil
	config.ExpectedStatusCode = 200
	config.WithHeader("Content-Type", "application/json")
	return config
}

func (config *ApiConfig) WithPath(path string) *ApiConfig {
	config.Path = path
	return config
}

func (config *ApiConfig) WithQuery(name string, value any) *ApiConfig {
	config.Query[name] = fmt.Sprintf("%v", value)
	return config
}

func (config *ApiConfig) WithHeader(name string, value any) *ApiConfig {
	config.Headers[name] = fmt.Sprintf("%v", value)
	return config
}

func (config *ApiConfig) WithBody(body any) *ApiConfig {
	config.Body = body
	return config
}

func (config *ApiConfig) WithExpectedStatusCode(expectedStatusCode int) *ApiConfig {
	config.ExpectedStatusCode = expectedStatusCode
	return config
}
