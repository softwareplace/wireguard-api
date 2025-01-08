package http_api

func (i *_impl[T]) Put(config *ApiConfig) (*T, error) {
	return i.exec("PUT", config)
}
