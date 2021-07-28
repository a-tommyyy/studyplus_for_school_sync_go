package fssync

import "encoding/json"

type Partner struct {
	PublicId       string `json:"public_id,omitempty"`
	CustomerUid    string `json:"customer_uid,omitempty"`
	SchoolName     string `json:"school_name"`
	Name           string `json:"name,omitempty"`
	TimeZone       string `json:"time_zone,omitempty"`
	ServerResponse `json:"-"`
}

type PartnerService struct {
	s *Service
}

type PartnerCreateCall struct {
	s       *Service
	partner *Partner
}

func NewPartnerService(s *Service) *PartnerService {
	return &PartnerService{s}
}

func (r *PartnerService) Create(partner *Partner) *PartnerCreateCall {
	c := &PartnerCreateCall{s: r.s}
	c.partner = partner
	return c
}

func (c *PartnerCreateCall) Do() (*Partner, error) {
	body, err := encodeJson(c.partner)
	if err != nil {
		return nil, err
	}
	res, err := request(c.s, "POST", "partners", body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	json.NewDecoder(res.Body).Decode(c.partner)
	c.partner.ServerResponse = ServerResponse{HTTPStatusCode: res.StatusCode}
	return c.partner, nil
}
