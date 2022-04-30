package maxdobleslicesum

func Solution(A []int) int {
	r := make([]int, len(A))
	l := make([]int, len(A))

	for i := 1; i < len(A)-1; i++ {
		max := l[i-1] + A[i]

		if max < 0 {
			max = 0
		}

		l[i] = max
	}

	// fmt.Println(l)

	for j := len(A) - 2; j > 0; j-- {
		max := r[j+1] + A[j]
		if max < 0 {
			max = 0
		}
		r[j] = max
	}

	// fmt.Println(r)

	maxDouble := 0
	for i := 1; i < len(A)-1; i++ {
		left := l[i-1]
		right := r[i+1]

		max := left + right
		// fmt.Println(i, left, right, max)
		if maxDouble < max {
			maxDouble = max
		}
	}

	return maxDouble
}
