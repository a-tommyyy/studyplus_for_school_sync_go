package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/atomiyama/studyplus_for_school_sync_go/fssync"
)

func main() {
	client := &http.Client{} // It should have injected Authorization Header when requesting.
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
