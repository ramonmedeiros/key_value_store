package keystore

import (
	"errors"
	"log/slog"
)

var (
	ErrNotFound      = errors.New("could not found key")
	ErrAlreadyExists = errors.New("key already exists")
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

func (k *KeyStore) AddKey(key string, value []byte) error {
	_, found := k.cache[key]
	if found {
		return ErrAlreadyExists
	}
	k.cache[key] = value
	return nil
}

func (k *KeyStore) GetKey(key string) ([]byte, error) {
	value, found := k.cache[key]
	if !found {
		return nil, ErrNotFound
	}
	return value, nil
}
