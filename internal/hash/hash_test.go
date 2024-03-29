package hash_test

import (
	"testing"

	"github.com/ramonmedeiros/key_value_store/internal/hash"
	"github.com/stretchr/testify/require"
)

func TestUsingHash(t *testing.T) {
	hashClient := hash.New()
	test1, err := hashClient.Get("1")
	require.NoError(t, err)

	test2, err := hashClient.Get("2")
	require.NoError(t, err)

	require.True(t, test1 > test2)
}
