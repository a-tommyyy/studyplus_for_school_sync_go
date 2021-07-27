package auth

import (
	"fmt"

	"golang.org/x/oauth2"
)

type Env string

const (
	EnvProduction  Env = "production"
	EnvDevelopment     = "development"
)

func EndpointFromEnv(env Env) (*oauth2.Endpoint, error) {
	switch env {
	case EnvProduction:
		return &oauth2.Endpoint{
			AuthURL:  "https://fs-lms.studyplus.co.jp/learning_material_supplier_api/v1/oauth/authorize",
			TokenURL: "https://fs-lms.studyplus.co.jp/learning_material_supplier_api/v1/oauth/token",
		}, nil
	case EnvDevelopment:
		return &oauth2.Endpoint{
			AuthURL:  "https://sandbox.fs-lms.studyplus.co.jp/learning_material_supplier_api/v1/oauth/authorize",
			TokenURL: "https://sandbox.fs-lms.studyplus.co.jp/learning_material_supplier_api/v1/oauth/token",
		}, nil
	default:
		return nil, fmt.Errorf("Invalid Env")
	}
}
