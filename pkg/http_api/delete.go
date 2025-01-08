package http_api

func (i *_impl[T]) Delete(config *ApiConfig) (*T, error) {
	return i.exec("DELETE", config)
}
