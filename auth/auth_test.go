package auth

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestNewAuthorization(t *testing.T) {
	app := mockApplication()
	actual := NewAuthorization(app)
	want := Authorization{
		Config: &oauth2.Config{
			ClientID:     app.ClientID,
			ClientSecret: app.ClientSecret,
			Scopes:       []string{"learning_material_supplier"},
			RedirectURL:  app.RedirectURL,
			Endpoint: oauth2.Endpoint{
				TokenURL: app.TokenURL,
				AuthURL:  app.AuthURL,
			},
		},
		Client: nil,
		ctx:    context.Background(),
	}

	assert.EqualValues(t, want, actual)
}

func TestAuthorizeURL(t *testing.T) {
	app := mockApplication()
	auth := NewAuthorization(app)
	authURL, _ := url.Parse(auth.AuthorizeURL())
	actualQuery := authURL.Query()
	wantQuery := url.Values{
		"client_id":     []string{app.ClientID},
		"response_type": []string{"code"},
		"redirect_uri":  []string{app.RedirectURL},
		"scope":         []string{"learning_material_supplier"},
		"state":         []string{"state"},
		"access_type":   []string{"offline"},
	}
	assert.EqualValues(t, wantQuery, actualQuery)
}

func mockApplication() *Application {
	return &Application{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		RedirectURL:  "REDIRECT_URL",
		TokenURL:     "https://example.com/oauth/token",
		AuthURL:      "https://example.com/oauth/authorize",
	}
}
