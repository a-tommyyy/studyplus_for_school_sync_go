package studyplus_for_school

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
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

func CreatePartner() (string, error) {
	baseUrl, err := url.Parse(os.Getenv("STUDYPLUS_FORSCHOOL_SYNC_BASE_URL"))
	fmt.Printf("baseurl: %s", baseUrl.Path)
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(baseUrl.Path, "/learning_material_supplier_api/v1/partners"), nil
}
