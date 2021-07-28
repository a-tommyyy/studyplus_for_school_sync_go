package fssync

type LearningMaterial struct {
	PublicId       string `json:"public_id"`
	CustmerUid     string `json:"customer_uid,omitempty"`
	Name           string `json:"name"`
	ImageUrl       string `json:"image_uri"`
	Unit           string `json:"unit"`
	ServerResponse `json:"-"`
}

type LearningMaterialService struct {
	s *Service
}

type LearningMaterialCreateCall struct {
	s                *Service
	learningMaterial *LearningMaterial
}

func (l *LearningMaterialService) Create(learningMaterial *LearningMaterial) *LearningMaterialCreateCall {
	c := LearningMaterialCreateCall{s: l.s}
	c.learningMaterial = learningMaterial
	return &c
}

func (c *LearningMaterialCreateCall) Do() (*LearningMaterial, error) {
	body, err := encodeJson(c.learningMaterial)
	if err != nil {
		return nil, err
	}
	resp, err := request(c.s, "POST", "learning_material_supplier_api/v1/learning_materials", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkResponse(resp); err != nil {
		return nil, err
	}
	learningMaterial := &LearningMaterial{
		ServerResponse: ServerResponse{HTTPStatusCode: resp.StatusCode},
	}
	target := &learningMaterial
	if err := decodeJson(&target, resp); err != nil {
		return nil, err
	}
	return learningMaterial, nil
}
