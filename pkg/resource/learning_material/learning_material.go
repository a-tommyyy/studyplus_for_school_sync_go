package learning_material

type LearningMaterial struct {
	PublicId   string `json:"public_id"`
	CustmerUid string `json:"customer_uid,omitempty"`
	Name       string `json:"name"`
	ImageUrl   string `json:"image_uri"`
	Unit       string `json:"unit"`
}
