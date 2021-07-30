package fssync

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullName(t *testing.T) {
	student := Student{FirstName: "FirstName", LastName: "LastName"}
	want := "LastName FirstName"
	assert.Equal(t, want, student.FullName())
}

func TestNewStudentService(t *testing.T) {
	s := &Service{}
	expected := &StudentService{s: s}
	assert.Equal(t, expected, NewStudentService(s))
}

func TestStudentService_Create(t *testing.T) {
	s, teardown := setupServer(t, &RequestAssertion{})
	defer teardown()
	student := &Student{}
	partnerID := "PartnerID"
	call := s.Partner.Student.Create(partnerID, student)
	expected := &StudentCreateCall{s: s, partnerID: partnerID, student: student}
	assert.Equal(t, expected, call)
}

func TestStudentCreateCall_Do(t *testing.T) {
	t.Run("received 401", func(t *testing.T) {
		student := &Student{
			FirstName:     "FirstName",
			LastName:      "LastName",
			FirstNameKana: "ファーストネーム",
			LastNameKana:  "ラストネーム",
		}
		partnerID := "partnerID"
		assertion := &RequestAssertion{
			description:         "it should raise error",
			expectedMethod:      "POST",
			expectedRequestPath: fmt.Sprintf("/partners/%s/students", partnerID),
			expectedRawQuery:    "",
			mockStatusCode:      401,
			mockResponseBody:    []byte(`{"messages": ["アクセストークンが無効です"]}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()
		_, err := s.Partner.Student.Create(partnerID, student).Do()
		assert.Error(t, err)
		assert.Equal(t, 401, student.ServerResponse.HTTPStatusCode)
	})

	t.Run("received 404 with invalid partnerID", func(t *testing.T) {
		student := &Student{
			FirstName:     "FirstName",
			LastName:      "LastName",
			FirstNameKana: "ファーストネーム",
			LastNameKana:  "ラストネーム",
		}
		partnerID := "InvalidPartnerID"
		assertion := &RequestAssertion{
			description:         "it should raise error",
			expectedMethod:      "POST",
			expectedRequestPath: fmt.Sprintf("/partners/%s/students", partnerID),
			expectedRawQuery:    "",
			mockStatusCode:      404,
			mockResponseBody:    []byte(`{"messages": ["Not found"]}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()
		_, err := s.Partner.Student.Create(partnerID, student).Do()
		assert.Error(t, err)
		assert.Equal(t, 404, student.ServerResponse.HTTPStatusCode)
	})

	t.Run("received 422", func(t *testing.T) {
		student := &Student{}
		partnerID := "PartnerID"
		assertion := &RequestAssertion{
			description:         "it should raise error",
			expectedMethod:      "POST",
			expectedRequestPath: fmt.Sprintf("/partners/%s/students", partnerID),
			expectedRawQuery:    "",
			mockStatusCode:      422,
			mockResponseBody: []byte(`{
				"messages": [
				"姓を入力してください",
				"名を入力してください",
				"カナ姓を入力してください",
				"カナ姓は全角カナで入力してください",
				"カナ名を入力してください",
				"カナ名は全角カナで入力してください"
				]
			}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()
		_, err := s.Partner.Student.Create(partnerID, student).Do()
		assert.Error(t, err)
		assert.Equal(t, 422, student.ServerResponse.HTTPStatusCode)
	})

	t.Run("received 200", func(t *testing.T) {
		partnerID := "PartnerID"
		student := &Student{
			FirstName:     "FirstName",
			LastName:      "LastName",
			FirstNameKana: "ファーストネーム",
			LastNameKana:  "ラストネーム",
		}
		studentPublicId := "samplePublicID"
		assertion := &RequestAssertion{
			description:         "it should return Student",
			expectedMethod:      "POST",
			expectedRequestPath: fmt.Sprintf("/partners/%s/students", partnerID),
			expectedRawQuery:    "",
			mockStatusCode:      200,
			mockResponseBody:    []byte(fmt.Sprintf(`{"public_id": "%s"}`, studentPublicId)),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()
		_, err := s.Partner.Student.Create(partnerID, student).Do()
		assert.NoError(t, err)
		assert.Equal(t, 200, student.ServerResponse.HTTPStatusCode)
		assert.Equal(t, studentPublicId, student.PublicId)
	})
}
