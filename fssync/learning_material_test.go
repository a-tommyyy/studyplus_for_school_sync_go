package fssync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLearningMaterialService(t *testing.T) {
	s := &Service{}
	expected := &LearningMaterialService{s: s}
	assert.Equal(t, expected, NewLearningMaterialService(s))
}

func TestLearningMaterialService_Create(t *testing.T) {
	s, teardown := setupServer(t, &RequestAssertion{})
	defer teardown()
	lm := &LearningMaterial{}
	call := s.LearningMaterial.Create(lm)
	expected := &LearningMaterialCreateCall{s: s, learningMaterial: lm}
	assert.Equal(t, expected, call)
}

func TestLearningMaterialCreateCall_Do(t *testing.T) {
	t.Run("received 401", func(t *testing.T) {
		lm := &LearningMaterial{
			Name:     "SampleName",
			ImageUrl: "https://example.com/sample.jpg",
			Unit:     "page",
		}
		assertion := &RequestAssertion{
			description:         "it should return error",
			expectedMethod:      "POST",
			expectedRequestPath: "/learning_materials",
			expectedRawQuery:    "",
			mockStatusCode:      401,
			mockResponseBody:    []byte(`{"messages": ["アクセストークンが無効です"]}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()

		_, err := s.LearningMaterial.Create(lm).Do()
		assert.Error(t, err)
	})

	t.Run("received 422", func(t *testing.T) {
		lm := &LearningMaterial{
			Name:     "",
			ImageUrl: "https://example.com/sample.jpg",
			Unit:     "page",
		}
		assertion := &RequestAssertion{
			description:         "it should return error",
			expectedMethod:      "POST",
			expectedRequestPath: "/learning_materials",
			expectedRawQuery:    "",
			mockStatusCode:      422,
			mockResponseBody:    []byte(`{"messages": ["Nameを入力してください"]}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()

		_, err := s.LearningMaterial.Create(lm).Do()
		assert.Error(t, err)
	})

	t.Run("received 200", func(t *testing.T) {
		lm := &LearningMaterial{
			Name:     "SampleName",
			ImageUrl: "https://example.com/sample.jpg",
		}
		assertion := &RequestAssertion{
			description:         "it should return LearningMaterial",
			expectedMethod:      "POST",
			expectedRequestPath: "/learning_materials",
			expectedRawQuery:    "",
			mockStatusCode:      200,
			mockResponseBody: []byte(`{
				"public_id": "82e83b0475",
				"name": "SampleName",
				"unit": "ページ",
				"image_url": "https://example.com/sample.jpg?timestamp=1627640769"
			}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()

		_, err := s.LearningMaterial.Create(lm).Do()
		assert.NoError(t, err)
		assert.Equal(t, 200, lm.ServerResponse.HTTPStatusCode)
	})
}
