package http_api

func (i *_impl[T]) Patch(config *ApiConfig) (*T, error) {
	return i.exec("PATCH", config)
}
