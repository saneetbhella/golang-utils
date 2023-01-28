package utils

import (
	"bytes"
	"net/http"
)

func MakeHttpRequest(
	client *http.Client,
	requestUrl string,
	method string,
	headers map[string]string,
	data []byte,
) (*http.Response, error) {
	req, _ := http.NewRequest(method, requestUrl, bytes.NewBuffer(data))

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
