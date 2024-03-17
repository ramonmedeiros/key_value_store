package keystore

import (
	"log/slog"
)

type KeyStore struct {
	logger *slog.Logger
	cache  map[string][]byte
}

//go:generate go run github.com/matryer/moq -pkg keystoretest -skip-ensure -out ./keystoretest/mock.go -stub . KeyStorer:Service
type KeyStorer interface {
	AddKey(key string, value []byte) error
	GetKey(key string) ([]byte, error)
}

func New(logger *slog.Logger) *KeyStore {
	return &KeyStore{
		logger: logger,
		cache:  map[string][]byte{},
	}
}

func (s *KeyStore) AddKey(key string, value []byte) error {
	return nil
}
func (s *KeyStore) GetKey(key string) ([]byte, error) {
	return nil, nil
}
