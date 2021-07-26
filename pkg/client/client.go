package client

import (
	"log"
	"net/url"
	"os"
	"time"
)

type Partner struct {
	customerUid string
	schoolName  string
	name        string
	timezone    time.Location
	startedOn   time.Time
}

// NOTE: 一旦URLを組み立てるところまで
func CreatePartner() string {
	return buildUrl("learning_material_supplier_api/v1/partners")
}

func buildUrl(p string) string {
	env := os.Getenv("STUDYPLUS_FOR_SCHOOL_SYNC_BASE_URL")
	baseUrl, err := url.Parse(env)
	if err != nil {
		log.Fatal(err)
	}

	rel, err := baseUrl.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	return rel.String()
}
