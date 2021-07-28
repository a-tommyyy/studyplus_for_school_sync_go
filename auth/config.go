package auth

import (
	"fmt"

	"golang.org/x/oauth2"
)

type Env string

const (
	EnvProduction  Env = "production"
	EnvDevelopment     = "development"
	EnvSandbox         = "sandbox"
)

func EndpointFromEnv(target *oauth2.Endpoint, env Env) error {
	switch env {
	case EnvProduction:
		target.AuthURL = "https://fs-lms.studyplus.co.jp/learning_material_supplier_api/v1/oauth/authorize"
		target.TokenURL = "https://fs-lms.studyplus.co.jp/learning_material_supplier_api/v1/oauth/token"
	case EnvDevelopment:
		target.AuthURL = "https://fs-lms.studyplus.co.jp.cage.boron.studylog.jp/learning_material_supplier_api/v1/oauth/authorize"
		target.TokenURL = "https://fs-lms.studyplus.co.jp.cage.boron.studylog.jp/learning_material_supplier_api/v1/oauth/token"
	case EnvSandbox:
		target.AuthURL = "https://sandbox.fs-lms.studyplus.co.jp/learning_material_supplier_api/v1/oauth/authorize"
		target.TokenURL = "https://sandbox.fs-lms.studyplus.co.jp/learning_material_supplier_api/v1/oauth/token"
	default:
		return fmt.Errorf("Invalid Env")
	}
	return nil
}
