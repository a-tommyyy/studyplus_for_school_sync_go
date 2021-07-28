package fssync

import (
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
	LearningMaterial            LearningMaterial
	Student                     Student
}
