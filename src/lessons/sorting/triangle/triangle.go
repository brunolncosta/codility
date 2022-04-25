package triangle

import "sort"

func Solution(A []int) int {
	sort.Ints(A)

	for i := 0; i <= len(A)-3; i++ {
		P := A[i]
		Q := A[i+1]
		R := A[i+2]

		if P+Q > R && Q+R > P && R+P > Q {
			return 1
		}
	}
	return 0
}
