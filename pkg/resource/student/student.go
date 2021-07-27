package student

import (
	"strings"
)

type Student struct {
	CustomerUid   string `json:"customer_uid"`
	PublicId      string `json:"public_id"`
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
