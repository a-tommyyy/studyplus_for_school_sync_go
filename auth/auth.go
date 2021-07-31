package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

type GetClientOptions map[string]string

func AuthorizeURL(cnf *oauth2.Config) string {
	return cnf.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func GetClientFromCode(ctx context.Context, cnf *oauth2.Config, authCode string) (*http.Client, error) {
	token, err := cnf.Exchange(ctx, authCode)
	if err != nil {
		return nil, err
	}
	return cnf.Client(ctx, token), nil
}

func GetClientFromFile(ctx context.Context, tokenFile string, cnf *oauth2.Config) (*http.Client, error) {
	f, err := os.Open(tokenFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var token *oauth2.Token
	json.NewDecoder(f).Decode(token)
	return cnf.Client(ctx, token), nil
}
