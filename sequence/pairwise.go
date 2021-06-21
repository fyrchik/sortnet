package sequence

// Bitonic returns swap sequence for pairwise sort.
func Pairwise(n int) SwapSequence {
	var ss SwapSequence

	a := 1
	for ; a < n; a *= 2 {
		var c int

		for b := a; b < n; {
			ss = append(ss, [2]int{b - a, b})

			b++
			c++
			if a <= c {
				c = 0
				b += a
			}
		}
	}

	a /= 4
	for e := 1; a > 0; a, e = a/2, e*2+1 {
		for d := e; d > 0; d /= 2 {
			var c int

			for b := (d + 1) * a; b < n; {
				ss = append(ss, [2]int{b - d*a, b})

				b++
				c++
				if a <= c {
					c = 0
					b += a
				}
			}
		}
	}

	return ss
}
