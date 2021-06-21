package sequence

// SwapSequence represents indices to perform compare-and-swap operations on.
type SwapSequence [][2]int

// Bitonic returns swap sequence for bitonic sort.
// There are at least 2 variants (only first is implemented):
// 1. Half-arrays are sorted in the same order and first step
//    of bitonic merge takes this into account (e.g. for 4 elements
//    it compares-and-swaps elements 0-3 and 1-2 instead of 0-2 and 1-3.
// 2. Half-arrays are sorted in the different order and bitonic
//    merge is performed as usual.
func Bitonic(n int) SwapSequence {
	var ss SwapSequence

	// 1. Sort consecutive pairs.
	for i := 1; i < n; i += 2 {
		ss = append(ss, [2]int{i - 1, i})
	}

	// 2. Merge sorted sub-arrays of increasing size.
	for groupSize := 2; groupSize < n; {
		groupSize *= 2
		for i := 0; i < n; i += groupSize {
			ss = bitonicMerge(ss, n, i, groupSize)
		}
	}

	return ss
}

func bitonicMerge(ss SwapSequence, n, start, sz int) SwapSequence {
	// 1. Form bitonic sequence out of 2 sorted ones.
	for l, r := start, start+sz-1; l < r; {
		if r < n {
			ss = append(ss, [2]int{l, r})
		}
		l++
		r--
	}

	// 2. Sort bitonic sequence.
	for gs := sz / 2; gs > 1; gs >>= 1 {
		for runStart := start; runStart < start+sz; runStart += gs {
			for i := runStart; i+gs/2 < runStart+gs; i++ {
				if i+gs/2 < n {
					ss = append(ss, [2]int{i, i + gs/2})
				}
			}
		}
	}

	return ss
}
