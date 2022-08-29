package sdkgen

import (
	"encoding/json"
	"errors"
	"os"
)

type TokenStoreInterface interface {
	Get() (AccessToken, error)
	Persist(token AccessToken) error
	Remove() error
}

type MemoryTokenStore struct {
	Token AccessToken
}

func (store MemoryTokenStore) Get() (AccessToken, error) {
	return store.Token, nil
}

func (store MemoryTokenStore) Persist(token AccessToken) error {
	store.Token = token
	return nil
}

func (store MemoryTokenStore) Remove() error {
	store.Token = AccessToken{}
	return nil
}

func NewMemoryTokenStore() MemoryTokenStore {
	return MemoryTokenStore{}
}

type FileTokenStore struct {
	path string
}

func (store FileTokenStore) Get() (AccessToken, error) {
	data, err := os.ReadFile(store.path)
	if err != nil {
		return AccessToken{}, errors.New("could not read Token store file")
	}

	var token AccessToken
	err = json.Unmarshal(data, &token)
	if err != nil {
		return AccessToken{}, errors.New("could not unmarshal access Token")
	}

	return token, nil
}

func (store FileTokenStore) Persist(token AccessToken) error {
	raw, err := json.Marshal(token)
	if err != nil {
		return err
	}

	err = os.WriteFile(store.path, raw, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (store FileTokenStore) Remove() error {
	err := os.Remove(store.path)
	if err != nil {
		return err
	}

	return nil
}

func NewFileTokenStore(path string) FileTokenStore {
	return FileTokenStore{path: path}
}
