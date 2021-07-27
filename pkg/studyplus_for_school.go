package studyplus_for_school

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

func HelloWorld() (string, error) {
	return fmt.Sprintf("Hello World"), nil
}

type Partner struct {
	customerUid string
	schoolName  string
	name        string
	timezone    time.Location
	startedOn   time.Time
}

func CreatePartner() string {
	baseUrl, err := url.Parse(os.Getenv("STUDYPLUS_FOR_SCHOOL_SYNC_BASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	url, err := baseUrl.Parse("learning_material_supplier_api/v1/partners")
	if err != nil {
		log.Fatal(err)
	}
	return url.String()
}
