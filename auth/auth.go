package auth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

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
