package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// PostHandlerFunc Post api handler
func PostHandlerFunc(url string, body interface{}, header *http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = *header
	client := http.Client{}
	return client.Do(request)
}

// GetHandlerFunc Get api handler
func GetHandlerFunc(url string, header *http.Header) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = *header
	client := http.Client{}
	return client.Do(request)
}

// DeleteHandleFunc Delet eapi handler
func DeleteHandleFunc(url string, header *http.Header) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = *header
	client := http.Client{}
	return client.Do(request)
}
