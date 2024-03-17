package keystore

import (
	"log/slog"
	"sync"
	"time"
)

type KeyStore struct {
	logger *slog.Logger
	cache  map[string][]byte
	mutex  mutex
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
		mutex:  mutex{&sync.RWMutex{}},
	}
}

// AddKey append a key/value and set it to be expired
func (k *KeyStore) AddKey(key string, value []byte) error {
	_, found := k.mutex.WithReadLock(func() ([]byte, bool) {
		value, found := k.cache[key]
		return value, found
	})
	if found {
		return ErrAlreadyExists
	}

	k.mutex.WithWriteLock(func() {
		k.cache[key] = value
	})

	go k.expireKey(key)
	return nil
}

// GetKey retrieves a value based on key
func (k *KeyStore) GetKey(key string) ([]byte, error) {
	value, found := k.mutex.WithReadLock(func() ([]byte, bool) {
		value, found := k.cache[key]
		return value, found
	})
	if !found {
		return nil, ErrNotFound
	}
	return value, nil
}

func (k *KeyStore) expireKey(key string) {
	time.AfterFunc(expireTime, func() {
		k.mutex.WithWriteLock(func() {
			delete(k.cache, key)
		})
	})
}
