package auth

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

// FileTokenStore is a data source managing tokens.
type FileTokenStore struct {
	Path string
}

func NewFileTokenStore(path string) *FileTokenStore {
	store := &FileTokenStore{Path: path}
	return store
}

func (s *FileTokenStore) Get() (*oauth2.Token, error) {
	f, err := os.Open(s.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	json.NewDecoder(f).Decode(token)
	return token, nil
}

func (s *FileTokenStore) Save(token *oauth2.Token) error {
	f, err := os.OpenFile(s.Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}
