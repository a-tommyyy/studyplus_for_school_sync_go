package client

import (
	"fmt"
	"os"
	"testing"
)

func TestCreatePartner(t *testing.T) {
	key := "STUDYPLUS_FOR_SCHOOL_SYNC_BASE_URL"
	host := "https://fs-lms.studyplus.co.jp"
	os.Setenv(key, host)
	actual := CreatePartner()

	rel := "learning_material_supplier_api/v1/partners"
	// NOTE: https://fs-lms.studyplus.co.jp/learning_material_supplier_api/v1/partners
	want := fmt.Sprintf("%s/%s", host, rel)
	if want != actual {
		t.Fatalf("expected %s but got %s", want, actual)
	}
}
