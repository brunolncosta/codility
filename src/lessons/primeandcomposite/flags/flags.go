package flags

func Solution(A []int) int {
	if len(A) < 3 {
		return 0
	}

	peaks := make([]int, 0)

	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
			i++
		}
	}

	if len(peaks) <= 0 {
		return 0
	}

	maxFlags := 1
	// for i := 2; i <= len(A); i++ {
	for i := 2; i*i <= len(A)+i; i++ {
		curr := peaks[0]
		count := 1
		for j := 1; j < len(peaks); j++ {
			distance := curr - peaks[j]
			if distance < 0 {
				distance = distance * -1
			}

			if distance >= i {
				count++
				curr = peaks[j]

				if count >= i {
					maxFlags = i
					break
				}
			}
		}

		// Efficiency
		//if maxFlags != i {
		//	break
		//}
	}

	return maxFlags
}
