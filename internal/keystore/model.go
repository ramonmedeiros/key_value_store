package keystore

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrNotFound      = errors.New("could not found key")
	ErrAlreadyExists = errors.New("key already exists")

	expireTime = time.Minute * 30
)

type mutex struct {
	*sync.RWMutex
}

func (m *mutex) WithReadLock(fn func() ([]byte, bool)) ([]byte, bool) {
	m.RLock()
	v, f := fn()
	m.RUnlock()
	return v, f
}

func (m *mutex) WithWriteLock(fn func()) {
	m.Lock()
	fn()
	m.Unlock()
}
