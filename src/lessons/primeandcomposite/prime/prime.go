package prime

func Solution(N int) int {

	if N < 2 {
		return 1
	}

	for i := 2; i+i <= N; i++ {
		if N%i == 0 {
			return 0
		}
	}

	return 1
}
