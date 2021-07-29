package fssync

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	t.Run("client is nil", func(t *testing.T) {
		_, err := NewService(nil, BaseURLDevelopment)
		assert.Error(t, err)
	})

	t.Run("it return Service", func(t *testing.T) {
		client := &http.Client{}
		s, err := NewService(client, BaseURLDevelopment)
		assert.NoError(t, err)
		assert.Equal(t, client, s.client)
		assert.Equal(t, BaseURLDevelopment, s.BaseURL)
	})
}

type RequestAssertion struct {
	// Description of test case
	description string

	// Mock response file path
	mockResponseBodyFile string

	// Expected request parameters
	expectedMethod      string
	expectedRequestPath string
	expectedRawQuery    string

	// Expected value Optional
	expectedHeaderOpts map[string][]string
}

func setupServer(t *testing.T, a *RequestAssertion) (*Service, func()) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		assert.Equal(t, a.expectedMethod, r.Method)
		assert.Equal(t, a.expectedRequestPath, r.URL.Path)
		assert.Equal(t, a.expectedRawQuery, r.URL.RawQuery)
		if a.expectedHeaderOpts != nil {
			for k, v := range a.expectedHeaderOpts {
				assert.Equal(t, v, r.Header[k])
			}
		}
		body, err := ioutil.ReadFile(a.mockResponseBodyFile)
		if err != nil {
			t.Fatalf("Failed to read mock response body file %s: %s\n", a.mockResponseBodyFile, err.Error())
		}
		rw.Write(body)
	}))
	s, err := NewService(server.Client(), BaseURL(server.URL))
	if err != nil {
		t.Fatal(err.Error())
	}
	fn := server.Close
	return s, fn
}
