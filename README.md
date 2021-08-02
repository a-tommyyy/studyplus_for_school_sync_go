# Studyplus for School SYNC API Client
# This project is currently alpha, possibility having breaking changes.
[![Maintainability](https://api.codeclimate.com/v1/badges/1704932645a35c702b79/maintainability)](https://codeclimate.com/github/atomiyama/studyplus_for_school_sync_go/maintainability)
[![github action workflow](https://github.com/atomiyama/studyplus_for_school_sync_go/actions/workflows/go.yml/badge.svg?branch:master)](https://github.com/atomiyama/studyplus_for_school_sync_go/actions?query=branch:master)

studyplus_for_school_sync_go is a API client for [Studyplus for School SYNC](https://studyplus.github.io/fs-sync-api/)

## Getting Started
### Authorization
**It is a NOT required implementation**, most users can use [golang.org/x/oauth2](https://pkg.go.dev/golang.org/x/oauth2) for Oauth2 authorization process.

```go
import (
	"fmt"
	"os"

	. "github.com/atomiyama/studyplus_for_school_sync_go/auth"
	"golang.org/x/oauth2"
)

type DBTokenStore struct {
	// some fields. 
}

func (s *DBTokenStore) Get() (*oauth2.Token, error) {
	// fetch persited token from DB.
}

func (s *DBTokenStore) Save(t *oauth2.Token) error {
	// Persist token.
}

func main() {
	var endpoint oauth2.Endpoint
	EndpointFromEnv(&endpoint, EnvDevelopment)
	cnf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"learning_material_supplier"},
		Endpoint:     endpoint,
	}
	src := &DBTokenStore{}
	authorization := NewAuthorization(cnf, src)
	url := authorization.AuthCodeURL("state")

	var code string // get auth code.

	ctx := context.Background()
	authorization.AuthorizeFromCode(ctx, code)

	ctx = context.Background()
	client, err := authorization.Client(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Request API via client.
}
```

### Creating LearningMaterial
```go
import (
	"fmt"
	"log"
	"net/http"

	. "github.com/atomiyama/studyplus_for_school_sync_go/fssync"
)

func main() {
	client := &http.Client{} // It should have injected Authorization Header when requesting.
	service, err := NewService(client, BaseURLDevelopment)
	if err != nil {
		log.Fatal(err)
	}
	lm := &LearningMaterial{
		Name:     "SAMPLE_NAME",
		ImageUrl: "https://exmaple.com/image.jpeg",
	}
	lm, err = service.LearningMaterial.Create(lm).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New LearningMaterial: %+v\n", lm)
}
```

### Creating Partner
TBD

### Creating Student
TBD

### Creating StudyRecord
TBD
