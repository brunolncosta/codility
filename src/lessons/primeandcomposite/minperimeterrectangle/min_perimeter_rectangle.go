package minperimeterrectangle

func Solution(N int) int {
	// write your code in Go 1.4
	min := -1
	i := 1
	for ; i*i <= N; i++ {
		if N%i == 0 {
			p := 2 * (i + (N / i))
			if min < 0 || p < min {
				min = p
			}
		}
	}

	return min
}
