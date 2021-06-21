package sequence

import (
	"math/bits"
)

// Batcher returns returns swap sequence for batcher odd-even sort.
func Batcher(n int) SwapSequence {
	return batcherSort(nil, n, 0)
}

func batcherSort(ss SwapSequence, n, start int) SwapSequence {
	if n-start <= 1 {
		return ss
	}

	// mid is a largest power of 2 smaller than n-start.
	mid := start + 1<<(bits.Len(uint(n-start-1))-1)

	ss = batcherSort(ss, mid, start)
	ss = batcherSort(ss, n, mid)
	ss = batcherMerge(ss, n, start, 1)

	return ss
}

func batcherMerge(ss SwapSequence, n, start, step int) SwapSequence {
	if n <= start+2*step {
		if start+step < n {
			ss = append(ss, [2]int{start, start + step})
		}
		return ss
	}

	ss = batcherMerge(ss, n, start, step*2)
	ss = batcherMerge(ss, n, start+step, step*2)

	for i := start + step; i+step < n; i += step * 2 {
		ss = append(ss, [2]int{i, i + step})
	}

	return ss
}
