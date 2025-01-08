package http_api

func (i *_impl[T]) Post(config *ApiConfig) (*T, error) {
	return i.exec("POST", config)
}
