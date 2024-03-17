package keystore

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

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

	time.Sleep(time.Second * 1)

	_, err = cache.GetKey(key)
	require.ErrorIs(t, ErrNotFound, err)
}
