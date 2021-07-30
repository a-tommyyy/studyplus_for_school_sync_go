package fssync

import (
	"encoding/json"
	"time"
)

type StudyRecord struct {
	RecordedAt                  time.Time `json:"recorded_at"`
	StartPosition               uint      `json:"start_position,omitempty"`
	EndPosition                 uint      `json:"end_position,omitemtpy"`
	Comment                     string    `json:"comment,omitemtpy"`
	ExternalLink                string    `json:"external_link,omitemtpy"`
	Amount                      uint      `json:"amount,omitemtpy"`
	NumberOfSeconds             uint      `json:"number_of_seconds,omitemtpy"`
	LearningMaterialPublicId    string    `json:"learning_material_public_id,omitemtpy"`
	StudentPublicId             string    `json:"student_public_id,omitemtpy"`
	LearningMaterialCustomerUid string    `json:"learning_material_customer_uid,omitemtpy"`
	StudentCustomerUid          string    `json:"student_customer_uid,omitemtpy"`
	ServerResponse              `json:"-"`
}

type StudyRecordService struct {
	s *Service
}

type StudyRecordCreateCall struct {
	s           *Service
	studyRecord *StudyRecord
}

type StudentOptions interface {
	Get() (key, value string)
}

func NewStudyRecordService(s *Service) *StudyRecordService {
	return &StudyRecordService{s: s}
}

func (r *StudyRecordService) Create(studyRecord *StudyRecord) *StudyRecordCreateCall {
	c := &StudyRecordCreateCall{s: r.s}
	c.studyRecord = studyRecord
	return c
}

func (c *StudyRecordCreateCall) Do() (*StudyRecord, error) {
	body, err := encodeJson(c.studyRecord)
	if err != nil {
		return nil, err
	}
	res, err := request(c.s, "POST", "study_records", body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	c.studyRecord.ServerResponse = ServerResponse{HTTPStatusCode: res.StatusCode}
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	json.NewDecoder(res.Body).Decode(c.studyRecord)
	return c.studyRecord, nil
}
