package prefixsums

// Solution The solution is based in prefix sum and counting technics.
// Mapping when a letter appears
func Solution(S string, P []int, Q []int) []int {
	N := len(S)
	M := len(P)

	// Mapping when a letter appears
	// propagating the last position it appears
	A := make([]int, N)
	C := make([]int, N)
	G := make([]int, N)
	T := make([]int, N)

	if S[0] != 'A' {
		A[0] = -1
	}

	if S[0] != 'C' {
		C[0] = -1
	}

	if S[0] != 'G' {
		G[0] = -1
	}

	if S[0] != 'T' {
		T[0] = -1
	}

	for i := 1; i < N; i++ {
		switch S[i] {
		case 'A':
			A[i] = i
			C[i] = C[i-1]
			G[i] = G[i-1]
			T[i] = T[i-1]
		case 'C':
			A[i] = A[i-1]
			C[i] = i
			G[i] = G[i-1]
			T[i] = T[i-1]
		case 'G':
			A[i] = A[i-1]
			C[i] = C[i-1]
			G[i] = i
			T[i] = T[i-1]
		case 'T':
			A[i] = A[i-1]
			C[i] = C[i-1]
			G[i] = G[i-1]
			T[i] = i
		}
	}

	result := make([]int, 0)

	for K := 0; K < M; K++ {
		if A[Q[K]] >= P[K] {
			result = append(result, 1)
			continue
		}

		if C[Q[K]] >= P[K] {
			result = append(result, 2)
			continue
		}

		if G[Q[K]] >= P[K] {
			result = append(result, 3)
			continue
		}

		if T[Q[K]] >= P[K] {
			result = append(result, 4)
			continue
		}
	}

	return result
}
