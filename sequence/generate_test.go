package sequence

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type sequenceFunc = func(int) SwapSequence

const maxTestSize = 16

func TestBitonic(t *testing.T) {
	for i := 1; i < maxTestSize; i++ {
		t.Run(strconv.FormatUint(uint64(i), 10), func(t *testing.T) {
			testGeneric(t, i, Bitonic)
		})
	}
}

func TestBatcher(t *testing.T) {
	for i := 1; i < maxTestSize; i++ {
		t.Run(strconv.FormatUint(uint64(i), 10), func(t *testing.T) {
			testGeneric(t, i, Batcher)
		})
	}
}

func TestPairwise(t *testing.T) {
	for i := 1; i < maxTestSize; i++ {
		t.Run(strconv.FormatUint(uint64(i), 10), func(t *testing.T) {
			testGeneric(t, i, Pairwise)
		})
	}
}

func testGeneric(t *testing.T, n int, f sequenceFunc) {
	ss := f(n)
	a := make([]int, n)

	// Test for all 0-1 sequences to ensure correctness in general case.
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			a[j] = (i >> j) & 1
		}
		b := cloneInts(a)
		sort.Ints(b)
		sortDirectly(a, ss)
		require.Equal(t, b, a)
	}
}

func sortDirectly(a []int, ss SwapSequence) {
	for _, xy := range ss {
		x, y := xy[0], xy[1]
		if a[x] > a[y] {
			a[x], a[y] = a[y], a[x]
		}
	}
}

func cloneInts(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}
