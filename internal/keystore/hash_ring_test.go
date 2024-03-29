package keystore

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/ramonmedeiros/key_value_store/internal/hash"
	"github.com/stretchr/testify/require"
)

// TestDistributionNodes prints amount of items per node
func TestDistributionNodes(t *testing.T) {
	expireTime = time.Second * 5
	keyStore, err := New(slog.New(slog.NewJSONHandler(os.Stderr, nil)), hash.New(), 5, 10)
	require.NoError(t, err)

	totalKeys := 100000

	for index := 0; index < totalKeys; index++ {
		key := "key_" + strconv.Itoa(index)
		value := []byte("value")
		err := keyStore.AddKey(key, value)
		require.NoError(t, err)
	}

	for _, node := range keyStore.nodes {
		fmt.Printf("%d items\n", len(node.cache))
	}

	for index := 0; index < totalKeys; index++ {
		key := "key_" + strconv.Itoa(index)
		value, err := keyStore.GetKey(key)
		require.NoError(t, err)
		require.Equal(t, "value", string(value))
	}
}
