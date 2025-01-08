package http_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (i *_impl[T]) build(method string, config *ApiConfig) (*http.Request, error) {
	var body io.Reader

	if config.Body != nil {
		jsonBody, err := json.Marshal(config.Body)
		if err != nil {
			log.Fatalf("Failed to marshal request body: %v", err)
		}
		body = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, config.Host+config.Path, body)

	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %v", err)
	}

	// Add headers to the request
	for key, value := range config.Headers {
		req.Header.Set(key, value)
	}

	// Add query parameters to the URL
	query := req.URL.Query()
	for key, value := range config.Query {
		query.Set(key, value)
	}

	return req, nil
}

func (i *_impl[T]) exec(method string, config *ApiConfig) (*T, error) {
	request, err := i.build(method, config)

	if err != nil {
		return nil, fmt.Errorf("failed to build request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("failed to make %s request: %v", request.Method, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != config.ExpectedStatusCode {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if &i.response != nil {
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&i.response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %v", err)
		}
	}

	return &i.response, nil
}
