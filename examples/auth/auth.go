package main

import (
	"fmt"
	"os"

	. "github.com/atomiyama/studyplus_for_school_sync_go/auth"
	"golang.org/x/oauth2"
)

func main() {
	var endpoint oauth2.Endpoint
	EndpointFromEnv(&endpoint, EnvDevelopment)
	cnf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"learning_material_supplier"},
		Endpoint:     endpoint,
	}
	src := &FileTokenStore{Path: "credentials.json"}
	authorization := NewAuthorization(cnf, src)
	authorization.AuthorizeCLI("state")
	fmt.Printf("Authorization succeeded. The token has been persisted.")
}
