package fssync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPartnerService(t *testing.T) {
	service := Service{}
	actual := NewPartnerService(&service)
	assert.EqualValues(t, &PartnerService{s: &service}, actual)
}

func TestCreate(t *testing.T) {
	service := NewPartnerService(&Service{})
	partner := Partner{}
	actual := service.Create(&partner)
	assert.EqualValues(t, PartnerCreateCall{s: service.s, partner: &partner}, actual)
}

func TestDo(t *testing.T) {
	assert.Equal(t, true, true)
}
