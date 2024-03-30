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
	"gonum.org/v1/gonum/stat"
)

// TestDistributionNodes prints amount of items per node
func TestDistributionNodes(t *testing.T) {
	testcases := []struct {
		name         string
		nodes        int
		virtualNodes int
		totalKeys    int
	}{
		{
			name:         "single node",
			nodes:        1,
			virtualNodes: 1,
			totalKeys:    100000,
		},
		{
			name:         "10 node, 1000 virtual",
			nodes:        10,
			virtualNodes: 1000,
			totalKeys:    1000000,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expireTime = time.Second * 5
			keyStore, err := New(slog.New(slog.NewJSONHandler(os.Stderr, nil)), hash.New(), tc.nodes, tc.virtualNodes)
			require.NoError(t, err)

			beforeAdd := time.Now()
			for index := 0; index < tc.totalKeys; index++ {
				key := "key_" + strconv.Itoa(index)
				value := []byte("value")
				err := keyStore.AddKey(key, value)
				require.NoError(t, err)
			}
			fmt.Printf("adding keys %s\n", time.Since(beforeAdd))

			beforeCalc := time.Now()
			pop := []float64{}
			weights := []float64{}
			for _, n := range keyStore.sortedNodes {
				pop = append(pop, float64(n))
				weights = append(weights, float64(len(keyStore.nodes[n].cache)))
			}

			standardDeviation := stat.StdDev(pop, weights)
			fmt.Printf("%s - The weighted standard deviation of the samples is %.4f\n", tc.name, standardDeviation)
			fmt.Printf("deviation %s\n", time.Since(beforeCalc))

			beforeGet := time.Now()
			for index := 0; index < tc.totalKeys; index++ {
				key := "key_" + strconv.Itoa(index)
				value, err := keyStore.GetKey(key)
				require.NoError(t, err)
				require.Equal(t, "value", string(value))
			}
			fmt.Printf("getting keys %s\n", time.Since(beforeGet))
		})
	}
}
