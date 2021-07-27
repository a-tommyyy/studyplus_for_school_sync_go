package studyplus_for_school

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	actual, err := HelloWorld()

	if err != nil {
		t.Fatalf("HelloWorld raised some error")
	}

	want := "Hello World"
	if want != actual {
		t.Fatalf("expected %s but got %s", want, actual)
	}
}

func TestCreatePartner(t *testing.T) {
	key := "STUDYPLUS_FOR_SCHOOL_SYNC_BASE_URL"
	value := "https://fs-lms.studyplus.co.jp"
	os.Setenv(key, value)
	actual := CreatePartner()
	want := "https://fs-lms.studyplus.co.jp/learning_material_supplier_api/v1/partners"
	assert.Equal(t, want, actual)
}
