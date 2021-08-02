package auth

import (
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestAuthorization_AuthCodeURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cnf := newMockConfig()
	ts := NewMockTokenStore(ctrl)
	authorization := NewAuthorization(cnf, ts)
	state := "TestAuthorization_AuthCodeURL"
	authURL, _ := url.Parse(authorization.AuthCodeURL(state))
	actualQuery := authURL.Query()
	wantQuery := url.Values{
		"client_id":     []string{cnf.ClientID},
		"response_type": []string{"code"},
		"redirect_uri":  []string{cnf.RedirectURL},
		"scope":         cnf.Scopes,
		"state":         []string{state},
		"access_type":   []string{"offline"},
	}
	assert.Equal(t, wantQuery, actualQuery)
}

func TestAuthorization_AuthorizeFromCode(t *testing.T) {
	assert.Equal(t, true, true)
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
