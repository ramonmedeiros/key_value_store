package keystore

import (
	"errors"
	"log/slog"
	"slices"
	"strconv"
	"sync"
	"time"

	"math/rand"

	"github.com/ramonmedeiros/key_value_store/internal/hash"
)

type KeyStore struct {
	logger      *slog.Logger
	nodes       map[uint32]node
	hashClient  hash.Hasher
	sortedNodes []uint32
}

type node struct {
	cache map[uint32][]byte
	mutex mutex
}

//go:generate go run github.com/matryer/moq -pkg keystoretest -with-resets -skip-ensure -out ./keystoretest/mock.go -stub . KeyStorer:Service
type KeyStorer interface {
	AddKey(key string, value []byte) error
	GetKey(key string) ([]byte, error)
}

func New(logger *slog.Logger, hashClient hash.Hasher, numberOfNodes int, virtualNodesPerNode int) (*KeyStore, error) {
	keyStore := &KeyStore{
		logger:     logger,
		hashClient: hashClient,
		nodes:      map[uint32]node{},
	}

	for n := 1; n <= numberOfNodes; n++ {
		for v := 1; v <= virtualNodesPerNode; v++ {
			nodeKey, err := hashClient.Get(strconv.Itoa(rand.Intn((10000 * n) + v)))
			if err != nil {
				return nil, err
			}

			keyStore.nodes[nodeKey] = node{
				cache: map[uint32][]byte{},
				mutex: mutex{&sync.RWMutex{}},
			}

			keyStore.sortedNodes = append(keyStore.sortedNodes, nodeKey)
		}
	}

	slices.Sort(keyStore.sortedNodes)

	return keyStore, nil
}

// AddKey append a key/value and set it to be expired
func (k *KeyStore) AddKey(key string, value []byte) error {
	hashKey, err := k.hashClient.Get(key)
	if err != nil {
		k.logger.Error("could not cast key", err)
		return err
	}

	nodeKey := k.findNodeKey(hashKey)
	n, ok := k.nodes[nodeKey]
	if !ok {
		return errors.New("could not find node")
	}

	_, found := n.mutex.WithReadLock(func() ([]byte, bool) {
		value, found := n.cache[hashKey]
		return value, found
	})
	if found {
		return ErrAlreadyExists
	}

	n.mutex.WithWriteLock(func() {
		n.cache[hashKey] = value
	})

	go k.expireKey(hashKey)
	return nil
}

// GetKey retrieves a value based on key
func (k *KeyStore) GetKey(key string) ([]byte, error) {
	hashKey, err := k.hashClient.Get(key)
	if err != nil {
		k.logger.Error("could not cast key", err)
		return nil, err
	}
	nodeKey := k.findNodeKey(hashKey)
	n, ok := k.nodes[nodeKey]
	if !ok {
		return nil, errors.New("could not find node")
	}

	value, found := n.mutex.WithReadLock(func() ([]byte, bool) {
		value, found := n.cache[hashKey]
		return value, found
	})
	if !found {
		return nil, ErrNotFound
	}
	return value, nil
}

// expireKey expires a key give a constant time
func (k *KeyStore) expireKey(key uint32) {
	time.AfterFunc(expireTime, func() {
		nodeKey := k.findNodeKey(key)
		n := k.nodes[nodeKey]
		n.mutex.WithWriteLock(func() {
			delete(k.nodes[nodeKey].cache, key)
		})
	})
}

// findNodeKey returns the index for node: defaults to first
func (k *KeyStore) findNodeKey(key uint32) uint32 {

	base := uint32(0)
	var nodeKey uint32
	for _, nodeKey = range k.sortedNodes {
		if key > base && key < nodeKey {
			break
		}
		base = nodeKey
	}
	return nodeKey
}
