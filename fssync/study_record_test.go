package fssync

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewStudyRecordService(t *testing.T) {
	s := &Service{}
	expected := &StudyRecordService{s: s}
	assert.Equal(t, expected, NewStudyRecordService(s))
}

func TestStudyRecordService_Create(t *testing.T) {
	s, teardown := setupServer(t, &RequestAssertion{})
	defer teardown()
	studyRecord := &StudyRecord{}
	call := s.StudyRecord.Create(studyRecord)
	expected := &StudyRecordCreateCall{s: s, studyRecord: studyRecord}
	assert.Equal(t, expected, call)
}

func TestStudyRecordCreateCall_Do(t *testing.T) {
	t.Run("received 401", func(t *testing.T) {
		studyRecord := &StudyRecord{
			StudentPublicId:          "StudentPublicID",
			LearningMaterialPublicId: "LearningMaterialPublicId",
			NumberOfSeconds:          3600,
			RecordedAt:               time.Now(),
		}
		assertion := &RequestAssertion{
			description:         "it should return error",
			expectedMethod:      "POST",
			expectedRequestPath: "/study_records",
			expectedRawQuery:    "",
			mockStatusCode:      401,
			mockResponseBody:    []byte(`{"messages": ["アクセストークンが無効です"]}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()

		_, err := s.StudyRecord.Create(studyRecord).Do()
		assert.Error(t, err)
		assert.Equal(t, 401, studyRecord.ServerResponse.HTTPStatusCode)
	})

	t.Run("received 404 with invalid studentID", func(t *testing.T) {
		studyRecord := &StudyRecord{
			StudentPublicId:          "",
			LearningMaterialPublicId: "LearningMaterialPublicId",
			NumberOfSeconds:          3600,
			RecordedAt:               time.Now(),
		}
		assertion := &RequestAssertion{
			description:         "it should return error",
			expectedMethod:      "POST",
			expectedRequestPath: "/study_records",
			expectedRawQuery:    "",
			mockStatusCode:      404,
			mockResponseBody:    []byte(`{"messages": ["Not found"]}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()

		_, err := s.StudyRecord.Create(studyRecord).Do()
		assert.Error(t, err)
		assert.Equal(t, 404, studyRecord.ServerResponse.HTTPStatusCode)
	})

	t.Run("received 404 with invalid LearningMaterialID", func(t *testing.T) {
		studyRecord := &StudyRecord{
			StudentPublicId:          "StudentPublicId",
			LearningMaterialPublicId: "",
			NumberOfSeconds:          3600,
		}
		assertion := &RequestAssertion{
			description:         "it should return error",
			expectedMethod:      "POST",
			expectedRequestPath: "/study_records",
			expectedRawQuery:    "",
			mockStatusCode:      404,
			mockResponseBody:    []byte(`{"messages": ["Not found"]}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()

		_, err := s.StudyRecord.Create(studyRecord).Do()
		assert.Error(t, err)
		assert.Equal(t, 404, studyRecord.ServerResponse.HTTPStatusCode)
	})

	t.Run("received 422", func(t *testing.T) {
		studyRecord := &StudyRecord{
			StudentPublicId:          "StudentPublicId",
			LearningMaterialPublicId: "LearningMaterialPublicId",
			NumberOfSeconds:          3600,
		}
		assertion := &RequestAssertion{
			description:         "it should return error",
			expectedMethod:      "POST",
			expectedRequestPath: "/study_records",
			expectedRawQuery:    "",
			mockStatusCode:      422,
			mockResponseBody: []byte(`{
				"messages": [
				"記録日時を入力してください",
				"記録日時は日付ではありません"
				]
			}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()

		_, err := s.StudyRecord.Create(studyRecord).Do()
		assert.Error(t, err)
		assert.Equal(t, 422, studyRecord.ServerResponse.HTTPStatusCode)
	})

	t.Run("received 200", func(t *testing.T) {
		studyRecord := &StudyRecord{
			StudentPublicId:          "StudentPublicId",
			LearningMaterialPublicId: "",
			NumberOfSeconds:          3600,
			RecordedAt:               time.Now(),
		}
		assertion := &RequestAssertion{
			description:         "it should return StudyRecord",
			expectedMethod:      "POST",
			expectedRequestPath: "/study_records",
			expectedRawQuery:    "",
			mockStatusCode:      200,
			mockResponseBody:    []byte(`{}`),
		}
		s, teardown := setupServer(t, assertion)
		defer teardown()

		_, err := s.StudyRecord.Create(studyRecord).Do()
		assert.NoError(t, err)
		assert.Equal(t, 200, studyRecord.ServerResponse.HTTPStatusCode)
	})
}
