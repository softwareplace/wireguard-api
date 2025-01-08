package http_api

func (i *_impl[T]) Get(config *ApiConfig) (*T, error) {
	return i.exec("GET", config)
}
