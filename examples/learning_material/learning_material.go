package main

import (
	"context"
	"fmt"
	"log"
	"os"

	. "github.com/atomiyama/studyplus_for_school_sync_go/auth"
	. "github.com/atomiyama/studyplus_for_school_sync_go/fssync"
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
	store := &FileTokenStore{Path: "credentials.json"}
	authorization := NewAuthorization(cnf, store)
	ctx := context.Background()
	client, err := authorization.Client(ctx)
	if err != nil {
		if err := authorization.AuthorizeCLI("state"); err != nil {
			log.Fatal(err)
		}
		client, err = authorization.Client(ctx)
	}

	service, err := NewService(client, BaseURLDevelopment)
	if err != nil {
		log.Fatal(err)
	}
	lm := &LearningMaterial{
		Name:     "SAMPLE_NAME",
		ImageUrl: "https://exmaple.com/image.jpeg",
	}
	lm, err = service.LearningMaterial.Create(lm).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New LearningMaterial: %+v\n", lm)
}
