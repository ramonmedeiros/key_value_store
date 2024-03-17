package keystore

import (
	"log/slog"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestExpire assert if one single key is expired
func TestExpire(t *testing.T) {
	expireTime = time.Second
	cache := New(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	key := "key"
	value := []byte("value")
	err := cache.AddKey(key, value)
	require.NoError(t, err)

	respValue, err := cache.GetKey(key)
	require.NoError(t, err)
	require.Equal(t, value, respValue)

	time.Sleep(expireTime + time.Second)

	respValue, err = cache.GetKey(key)
	require.Nil(t, respValue)
	require.ErrorIs(t, err, ErrNotFound)
}

// TestExpirationWithLotsofKeys assert if several keys are expired
func TestExpirationWithLotsofKeys(t *testing.T) {
	expireTime = time.Second * 5
	cache := New(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	totalKeys := 10000

	for index := 0; index < totalKeys; index++ {
		key := "key_" + strconv.Itoa(index)
		value := []byte("value")
		err := cache.AddKey(key, value)
		require.NoError(t, err)
	}

	time.Sleep(expireTime + time.Second)

	for index := 0; index < totalKeys; index++ {
		key := "key_" + strconv.Itoa(index)
		_, err := cache.GetKey(key)
		require.ErrorIs(t, ErrNotFound, err)
	}
}
