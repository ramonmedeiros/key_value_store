package server

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/ramonmedeiros/key_value_store/internal/keystore"
	"github.com/ramonmedeiros/key_value_store/internal/keystore/keystoretest"
	"github.com/stretchr/testify/require"
)

// TestAddGetKeys check all errors expected from test
func TestAddGetKeys(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	keystoreMock := keystoretest.Service{
		AddKeyFunc: func(key string, value []byte) error {
			return nil
		},
		GetKeyFunc: func(key string) ([]byte, error) {
			return []byte("test"), nil
		},
	}
	rest := New("8080", logger, &keystoreMock)

	t.Run("getKeyMethod", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		require.NoError(t, err)
		rest.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, "test", w.Body.String())
		require.Len(t, keystoreMock.GetKeyCalls(), 1)
		keystoreMock.ResetCalls()
	})

	t.Run("getKey not found", func(t *testing.T) {
		keystoreMock.GetKeyFunc = func(key string) ([]byte, error) {
			return nil, keystore.ErrNotFound
		}
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/test", nil)
		require.NoError(t, err)
		rest.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusNotFound, w.Code)
		require.Len(t, keystoreMock.GetKeyCalls(), 1)
		keystoreMock.ResetCalls()
	})

	t.Run("addKeyMethod", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/test", io.NopCloser(strings.NewReader("test")))
		require.NoError(t, err)
		rest.router.ServeHTTP(w, req)

		require.Equal(t, http.StatusCreated, w.Code)
		require.Len(t, keystoreMock.AddKeyCalls(), 1)
		require.Equal(t, []byte("test"), keystoreMock.AddKeyCalls()[0].Value)
	})
}
