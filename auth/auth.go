package auth

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type Authorization struct {
	Config *oauth2.Config
	Client *http.Client
	ctx    context.Context
}
type Application struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	TokenURL     string
	AuthURL      string
}

/*
	EXAMPLE IMPLEMENTATION
	func main() {
		app := &Application{
			ClientID: os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			RedirectURL: os.Getenv("REDIRECT_URL"),
			TokenURL: os.Getenv("TOKEN_URL"),
			AuthURL: os.Getenv("AUTH_URL"),
		}
		auth := NewAuthorization(app)
		fmt.Printf("Visit the URL for the auth dialog:\n%v\n", auth.AuthorizeURL)

		var code string
		if _, err := fmt.Scan(&code); err != nil {
			log.Fatal(err)
		}

		auth.AuthrizeFromCode(code)
	}
*/
func NewAuthorization(app *Application) Authorization {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     app.ClientID,
		ClientSecret: app.ClientSecret,
		Scopes:       []string{"learning_material_supplier"},
		RedirectURL:  app.RedirectURL,
		Endpoint: oauth2.Endpoint{
			TokenURL: app.TokenURL,
			AuthURL:  app.AuthURL,
		},
	}
	auth := Authorization{Config: conf, Client: nil, ctx: ctx}
	return auth
}

func (auth *Authorization) AuthorizeURL() string {
	return auth.Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (auth *Authorization) AuthorizeFromCode(code string) *Authorization {
	auth.ctx = context.WithValue(auth.ctx, oauth2.HTTPClient, &http.DefaultClient)
	token, err := auth.Config.Exchange(auth.ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := auth.Config.Client(auth.ctx, token)
	auth.Client = client
	return auth
}
