package maxslice

func Solution(A []int) int {
	if len(A) <= 1 {
		return 0
	}

	profits := make([]int, 0)
	lastPrice := A[0]
	for i := 1; i < len(A); i++ {
		profit := A[i] - lastPrice
		profits = append(profits, profit)
		lastPrice = A[i]
	}

	maxEnd := 0
	maxSlice := 0
	for _, profit := range profits {
		maxEnd = maxEnd + profit
		if maxEnd < 0 {
			maxEnd = 0
		}
		if maxSlice < maxEnd {
			maxSlice = maxEnd
		}
	}
	return maxSlice
}
