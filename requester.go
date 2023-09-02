package randommer

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type requester struct {
	BaseUrl string
	ApiKey  string
	http.Client
}

func (r *requester) buildPath(path string) string {
	fullPath := fmt.Sprintf("%s/%s", r.BaseUrl, path)
	log.Println(fullPath)
	return fullPath
}

func (r *requester) Get(path string) (*http.Response, error) {
	new_request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	new_request.Header.Set("X-API-KEY", r.ApiKey)
	new_request.Header.Set("Content-Type", "application/json")
	res, err := r.Do(new_request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (r *requester) Post(path string, payload_string string) (*http.Response, error) {
	data := []byte(payload_string)
	new_request, _ := http.NewRequest("POST", path, bytes.NewBuffer(data))
	new_request.Header.Set("X-Api-Key", r.ApiKey)
	new_request.Header.Set("Content-Type", "application/json")
	res, err := r.Do(new_request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil

}

func newRequester(baseUrl, apiKey string) *requester {
	return &requester{BaseUrl: baseUrl, ApiKey: apiKey}
}
