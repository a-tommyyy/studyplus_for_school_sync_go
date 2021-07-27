package fssync

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Service struct {
	Client  *http.Client
	BaseURL *url.URL
}

func NewService(client *http.Client, base string) *Service {
	baseUrl, err := url.Parse(base)
	if err != nil {
		log.Fatal(err)
	}
	return &Service{client, baseUrl}
}

func request(s *Service, method string, path string, body io.Reader) (*http.Response, error) {
	endpoint := buildUrl(s, path)
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	return s.Client.Do(req)
}

type ServerResponse struct {
	HTTPStatusCode int
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}
	return fmt.Errorf("[STATUS: %v]", resp.StatusCode)
}

func encodeJson(v interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)
	err := json.NewDecoder(buf).Decode(v)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func decodeJson(target interface{}, res *http.Response) error {
	if res.StatusCode != http.StatusNoContent {
		return nil
	}
	return json.NewDecoder(res.Body).Decode(target)
}

func buildUrl(s *Service, p string) string {
	url, err := s.BaseURL.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	return url.String()
}
