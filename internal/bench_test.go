package internal

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"

	"github.com/fyrchik/sortnet/internal/networks"
)

func BenchmarkSort(b *testing.B) {
	b.ReportAllocs()
	for _, n := range []int{4, 8, 16} {
		b.Run(strconv.FormatUint(uint64(n), 10), func(b *testing.B) {
			benchmarkSort(b, n)
		})
	}
}

func benchmarkSort(b *testing.B, n int) {
	a := make([]int, n)
	for i := range a {
		a[i] = rand.Int()
	}

	const count = 1000

	b.Run("stdlib", func(b *testing.B) {
		curr := cloneInts(a)
		for i := 0; i < b.N; i++ {
			for j := 0; j < count; j++ {
				copy(curr, a)
				sort.Sort(sort.IntSlice(curr))
			}
		}
	})

	b.Run("bitonic", func(b *testing.B) {
		curr := cloneInts(a)
		f := getSortingFunc(n, "bitonic")
		for i := 0; i < b.N; i++ {
			for j := 0; j < count; j++ {
				copy(curr, a)
				f(sort.IntSlice(curr))
			}
		}
	})

	b.Run("batcher", func(b *testing.B) {
		curr := cloneInts(a)
		f := getSortingFunc(n, "batcher")
		for i := 0; i < b.N; i++ {
			for j := 0; j < count; j++ {
				copy(curr, a)
				f(sort.IntSlice(curr))
			}
		}
	})
}

func getSortingFunc(n int, name string) func(sort.Interface) {
	switch name {
	case "batcher":
		switch n {
		case 4:
			return networks.Batcher4
		case 8:
			return networks.Batcher8
		case 16:
			return networks.Batcher16
		}
	case "bitonic":
		switch n {
		case 4:
			return networks.Bitonic4
		case 8:
			return networks.Bitonic8
		case 16:
			return networks.Bitonic16
		}
	}
	panic("unexpected")
}

func cloneInts(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}
