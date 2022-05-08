package count_factors

func Solution(N int) int {
	factors := 0
	i := 1
	for ; i*i < N; i++ {
		if N%i == 0 {
			factors = factors + 2
		}
	}

	if i*i == N {
		factors++
	}

	return factors
}
