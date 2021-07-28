package fssync

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
)

const (
	BaseURLProduction  = "https://fs-lms.studyplus.co.jp/learning_material_supplier_api/v1"
	BaseURLSandbox     = "https://sandbox.fs-lms.studyplus.co.jp/learning_material_supplier_api/v1"
	BaseURLDevelopment = "https://fs-lms.studyplus.co.jp.cage.boron.studylog.jp/learning_material_supplier_api/v1"
)

type Service struct {
	client           *http.Client
	BaseURL          string
	LearningMaterial *LearningMaterialService
}

func NewService(client *http.Client, baseURL string) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BaseURL: baseURL}
	s.LearningMaterial = NewLearningMaterialService(s)
	return s, nil
}

func request(s *Service, method string, path string, body io.Reader) (*http.Response, error) {
	endpoint := buildUrl(s, path)
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return s.client.Do(req)
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
	bodyBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(bodyBytes), nil
}

func buildUrl(s *Service, p string) string {
	url, err := url.Parse(s.BaseURL)
	if err != nil {
		log.Fatal(err)
	}
	url.Path = path.Join(url.Path, p)
	return url.String()
}
