package countsemiprimes

func Solution(N int, P []int, Q []int) []int {
	sieve := make([]bool, N+1)

	for i := 2; i*i <= N; i++ {
		if !sieve[i] {
			for j := i * i; j < N; j += i {
				sieve[j] = true
			}
		}
	}

	sieveSemi := make([]bool, N+1)

	for i := 2; i*i <= N; i++ {
		if !sieve[i] {
			for j := i; i*j <= N; j++ {
				if !sieve[j] {
					sieveSemi[i*j] = true
				}
			}
		}
	}

	count := 0
	prefixSum := make([]int, N+1)
	semi := make([]int, 0)
	for i, s := range sieveSemi {
		if s {
			semi = append(semi, i)
			count++
		}
		prefixSum[i] = count
	}

	// fmt.Println(semi)
	// fmt.Println(prefixSum)

	result := make([]int, 0)
	for i := 0; i < len(P); i++ {
		count := prefixSum[Q[i]] - prefixSum[P[i]-1]
		result = append(result, count)
	}

	return result
}
