package resource

import (
	"strings"
	"time"
)

type LearningMaterial struct {
	PublicId   string `json:"public_id"`
	CustmerUid string `json:"customer_uid,omitempty"`
	Name       string `json:"name"`
	ImageUrl   string `json:"image_uri"`
	Unit       string `json:"unit"`
}

type Partner struct {
	PublicId    string `json:"public_id"`
	CustomerUid string `json:"customer_uid,omitempty"`
	SchoolName  string `json:"school_name"`
	Name        string `json:"name"`
	TimeZone    string `json:"time_zone"`
}

type Student struct {
	PublicId      string `json:"public_id"`
	CustomerUid   string `json:"customer_uid,omitempty"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	FirstNameKana string `json:"first_name_kana"`
	LastNameKana  string `json:"last_name_kana"`
	SchoolType    Grade  `json:"school_type"`
	Grade         uint8  `json:"grade"`
	Code          string `json:"code"`
}

type Grade uint8

const (
	Unknown Grade = iota
	JuniorHighSchool
	HighSchool
	TechnicalCollege
	University
	GraduateSchool
	JuniorCollege
	CarrerCollege
	HighSchoolGraduate
	Working
	ElementarySchool
)

func (student *Student) FullName() string {
	return strings.Join([]string{student.LastName, student.FirstName}, " ")
}

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
