package http_api

func (i *_impl[T]) Head(config *ApiConfig) (*T, error) {
	return i.exec("HEAD", config)
}
