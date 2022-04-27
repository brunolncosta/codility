package maxslice

func Solution(A []int) int {
	maxEnd := 0
	maxSlice := 0
	hasNegative := false
	minNegative := -2147483648
	for _, a := range A {

		if a < 0 && a > minNegative {
			hasNegative = true
			minNegative = a
		}

		maxEnd = maxEnd + a
		if maxEnd < 0 {
			maxEnd = 0
		}

		if maxEnd > maxSlice {
			maxSlice = maxEnd
		}
	}

	if maxSlice == 0 && hasNegative {
		return minNegative
	}

	return maxSlice
}
