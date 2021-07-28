package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/atomiyama/studyplus_for_school_sync_go/auth"
	"github.com/atomiyama/studyplus_for_school_sync_go/fssync"
	"golang.org/x/oauth2"
)

func main() {
	var endpoint oauth2.Endpoint
	auth.EndpointFromEnv(&endpoint, auth.EnvDevelopment)
	cnf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"learning_material_supplier"},
		Endpoint:     endpoint,
	}
	url := auth.AuthorizeURL(cnf)
	fmt.Printf("Visit the URL:\n%v\n", url)
	fmt.Printf("Enter AuthCode>>")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client, err := auth.GetClientFromCode(ctx, cnf, code)
	if err != nil {
		log.Fatal(err)
	}

	service, err := fssync.NewService(client, fssync.BaseURLDevelopment)
	if err != nil {
		log.Fatal(err)
	}
	lm := &fssync.LearningMaterial{
		Name:     "SAMPLE_NAME",
		ImageUrl: "https://exmaple.com/image.jpeg",
	}
	lm, err = service.LearningMaterial.Create(lm).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New LearningMaterial: %+v", lm)
}
