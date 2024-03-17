package keystore

import (
	"errors"
	"log/slog"
	"sync"
	"time"
)

var (
	ErrNotFound      = errors.New("could not found key")
	ErrAlreadyExists = errors.New("key already exists")

	expireTime = time.Minute * 30
)

type KeyStore struct {
	logger *slog.Logger
	cache  map[string][]byte
	mutex  *sync.RWMutex
}

//go:generate go run github.com/matryer/moq -pkg keystoretest -with-resets -skip-ensure -out ./keystoretest/mock.go -stub . KeyStorer:Service
type KeyStorer interface {
	AddKey(key string, value []byte) error
	GetKey(key string) ([]byte, error)
}

func New(logger *slog.Logger) *KeyStore {
	return &KeyStore{
		logger: logger,
		cache:  map[string][]byte{},
		mutex:  &sync.RWMutex{},
	}
}

func (k *KeyStore) AddKey(key string, value []byte) error {
	k.mutex.RLock()
	_, found := k.cache[key]
	k.mutex.RUnlock()
	if found {
		return ErrAlreadyExists
	}
	k.mutex.Lock()
	k.cache[key] = value
	k.mutex.Unlock()

	go k.expireKey(key)
	return nil
}

func (k *KeyStore) GetKey(key string) ([]byte, error) {
	k.mutex.RLock()
	value, found := k.cache[key]
	k.mutex.RUnlock()
	if !found {
		return nil, ErrNotFound
	}
	return value, nil
}

func (k *KeyStore) expireKey(key string) {
	time.AfterFunc(expireTime, func() {
		k.mutex.RLock()
		delete(k.cache, key)
		k.mutex.RUnlock()
	})
}
