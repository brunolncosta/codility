package peaks

func Solution(A []int) int {
	if len(A) < 3 {
		return 0
	}

	peaks := make([]int, 0)
	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}

	if len(peaks) < 1 {
		return 0
	}

	max := 1
	for i := 2; i <= len(peaks); i++ {
		// could not divide to equals elements
		if len(A)%i != 0 {
			continue
		}

		y := 0
		init := 0
		end := (len(A) / i) - 1
		hasPeak := true
		for j := 0; j < i; j++ {
			hasBlockPeak := false
			for ; y < len(peaks); y++ {
				if peaks[y] >= init && peaks[y] <= end {
					init = end + 1
					end = end + (len(A))/i
					hasBlockPeak = true
					break
				}
			}

			if !hasBlockPeak {
				hasPeak = false
				break
			}
		}

		if hasPeak {
			max = i
		}
	}
	return max
}
