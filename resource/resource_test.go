package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullName(t *testing.T) {
	student := Student{FirstName: "FirstName", LastName: "LastName"}
	want := "LastName FirstName"
	assert.Equal(t, want, student.FullName())
}
