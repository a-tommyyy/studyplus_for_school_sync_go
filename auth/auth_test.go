package auth

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestAuthorizeURL(t *testing.T) {
	cnf := newMockConfig()
	authURL, _ := url.Parse(AuthorizeURL(cnf))
	actualQuery := authURL.Query()
	wantQuery := url.Values{
		"client_id":     []string{cnf.ClientID},
		"response_type": []string{"code"},
		"redirect_uri":  []string{cnf.RedirectURL},
		"scope":         cnf.Scopes,
		"state":         []string{"state"},
		"access_type":   []string{"offline"},
	}
	assert.EqualValues(t, wantQuery, actualQuery)
}

func newMockConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		RedirectURL:  "REDIRECT_URL",
		Scopes:       []string{"learning_material_supplier"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://example.com/authorize",
			TokenURL: "https://example.com/token",
		},
	}
}
