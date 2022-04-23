package maxproductofthree

// you can also use imports, for example:
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)

	if N < 3 {
		return 0
	}

	sort.Ints(A)
	max := A[N-1] * A[N-2] * A[N-3]

	maxWithNegative := A[0] * A[1] * A[N-1]
	if maxWithNegative > max {
		return maxWithNegative
	}

	return max
}
