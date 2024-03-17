package keystore_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/ramonmedeiros/key_value_store/internal/keystore"
	"github.com/stretchr/testify/require"
)

func TestAddGetKeys(t *testing.T) {
	testcases := []struct {
		name  string
		key   string
		value []byte
	}{
		{
			name:  "english string",
			key:   "key",
			value: []byte("value"),
		},
		{
			name:  "greek string",
			key:   "κλειδί",
			value: []byte("αξία"),
		},
		{
			name:  "japanese string",
			key:   "鍵",
			value: []byte("価値"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			cache := keystore.New(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

			value, err := cache.GetKey(tc.key)
			require.ErrorIs(t, keystore.ErrNotFound, err)
			require.Nil(t, value)

			err = cache.AddKey(tc.key, tc.value)
			require.NoError(t, err)

			value, err = cache.GetKey(tc.key)
			require.NoError(t, err)
			require.Equal(t, tc.value, value)

			err = cache.AddKey(tc.key, tc.value)
			require.ErrorIs(t, keystore.ErrAlreadyExists, err)
		})
	}
}
