package fssync

import (
	"encoding/json"
)

type LearningMaterial struct {
	PublicId       string `json:"public_id,omitempty"`
	CustmerUid     string `json:"customer_uid,omitempty"`
	Name           string `json:"name"`
	ImageUrl       string `json:"image_uri,omitempty"`
	Unit           string `json:"unit,omitempty"`
	ServerResponse `json:"-"`
}

type LearningMaterialService struct {
	s *Service
}

type LearningMaterialCreateCall struct {
	s                *Service
	learningMaterial *LearningMaterial
}

func NewLearningMaterialService(s *Service) *LearningMaterialService {
	return &LearningMaterialService{s}
}

func (r *LearningMaterialService) Create(learningMaterial *LearningMaterial) *LearningMaterialCreateCall {
	c := LearningMaterialCreateCall{s: r.s}
	c.learningMaterial = learningMaterial
	return &c
}

func (c *LearningMaterialCreateCall) Do() (*LearningMaterial, error) {
	body, err := encodeJson(c.learningMaterial)
	if err != nil {
		return nil, err
	}
	resp, err := request(c.s, "POST", "learning_materials", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkResponse(resp); err != nil {
		return nil, err
	}
	json.NewDecoder(resp.Body).Decode(c.learningMaterial)
	c.learningMaterial.ServerResponse = ServerResponse{HTTPStatusCode: resp.StatusCode}
	return c.learningMaterial, nil
}
