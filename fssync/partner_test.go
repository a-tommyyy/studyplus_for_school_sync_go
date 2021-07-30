package fssync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPartnerService(t *testing.T) {
	s := &Service{}
	actual := NewPartnerService(s)
	expected := &PartnerService{s: s, Student: &StudentService{s: s}}
	assert.EqualValues(t, expected, actual)
}

func TestPartnerService_Create(t *testing.T) {
	s, teardown := setupServer(t, &RequestAssertion{})
	defer teardown()
	partner := &Partner{}
	call := s.Partner.Create(partner)
	expected := &PartnerCreateCall{s: s, partner: partner}
	assert.EqualValues(t, expected, call)
}

func TestPartnerCreateCall_Do(t *testing.T) {
	t.Run("received 401", func(t *testing.T) {
		partner := &Partner{
			CustomerUid: "SampleCustomerUid",
			SchoolName:  "SampleSchoolName",
		}
		assertion := &RequestAssertion{
			description:         "it shold return error",
			expectedMethod:      "POST",
			expectedRequestPath: "/partners",
			expectedRawQuery:    "",
			mockStatusCode:      401,
			mockResponseBody: []byte(`{
				"messages": [
				"アクセストークンが無効です"
				]
			}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()

		_, err := s.Partner.Create(partner).Do()
		assert.Error(t, err)
	})

	t.Run("received 422", func(t *testing.T) {
		partner := &Partner{
			CustomerUid: "SampleCustomerUid",
			SchoolName:  "",
		}
		assertion := &RequestAssertion{
			description:         "it shold return error",
			expectedMethod:      "POST",
			expectedRequestPath: "/partners",
			expectedRawQuery:    "",
			mockStatusCode:      422,
			mockResponseBody:    []byte(`{"messages": ["School_nameを入力してください"]}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()

		_, err := s.Partner.Create(partner).Do()
		assert.Error(t, err)
	})

	t.Run("receved 200", func(t *testing.T) {
		partner := &Partner{
			CustomerUid: "SampleCustomerUid",
			SchoolName:  "SampleSchoolName",
		}
		assertion := &RequestAssertion{
			description:         "it shold return Partner",
			expectedMethod:      "POST",
			expectedRequestPath: "/partners",
			expectedRawQuery:    "",
			mockStatusCode:      200,
			mockResponseBody:    []byte(`{"public_id": "e5abde4f78"}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()

		_, err := s.Partner.Create(partner).Do()
		assert.NoError(t, err)
		assert.Equal(t, 200, partner.ServerResponse.HTTPStatusCode)
	})
}
