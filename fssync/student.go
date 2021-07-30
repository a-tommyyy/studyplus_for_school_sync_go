package fssync

import (
	"encoding/json"
	"strings"
)

type Student struct {
	PublicId       string     `json:"public_id,omitempty"`
	CustomerUid    string     `json:"customer_uid,omitempty"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	FirstNameKana  string     `json:"first_name_kana"`
	LastNameKana   string     `json:"last_name_kana"`
	SchoolType     SchoolType `json:"school_type,omitempty"`
	Grade          uint       `json:"grade,omitempty"`
	Code           string     `json:"code,omitempty"`
	ServerResponse `json:"-"`
}

type SchoolType uint

const (
	Unknown            SchoolType = 0
	JuniorHighSchool              = 1
	HighSchool                    = 2
	TechnicalCollege              = 3
	University                    = 4
	GraduateSchool                = 5
	JuniorCollege                 = 6
	CarrerCollege                 = 7
	HighSchoolGraduate            = 8
	Working                       = 9
	ElementarySchool              = 10
)

func (student *Student) FullName() string {
	return strings.Join([]string{student.LastName, student.FirstName}, " ")
}

type StudentService struct {
	s *Service
}

type StudentCreateCall struct {
	s         *Service
	student   *Student
	partnerID string
}

func NewStudentService(s *Service) *StudentService {
	return &StudentService{s: s}
}

func (r *StudentService) Create(partnerID string, student *Student) *StudentCreateCall {
	c := &StudentCreateCall{s: r.s}
	c.partnerID = partnerID
	c.student = student
	return c
}

func (c *StudentCreateCall) Do() (*Student, error) {
	body, err := encodeJson(c.student)
	if err != nil {
		return nil, err
	}
	url := "partners/" + c.partnerID + "/students"
	res, err := request(c.s, "POST", url, body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	c.student.ServerResponse = ServerResponse{HTTPStatusCode: res.StatusCode}
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	json.NewDecoder(res.Body).Decode(c.student)
	return c.student, nil
}
