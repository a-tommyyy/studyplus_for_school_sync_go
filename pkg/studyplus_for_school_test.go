package studyplus_for_school

import (
	"testing"
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
	actual, err := CreatePartner()

	if err != nil {
		t.Fatalf("CreatePartner raised some error")
	}

	want := "https://fs-lms.studyplus.co.jp/learning_material_supplier_api/v1/partners"
	if want != actual {
		t.Fatalf("expected %s but got %s", want, actual)
	}
}
