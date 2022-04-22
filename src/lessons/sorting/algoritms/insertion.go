package algoritms

// InsertionSort Very simple sort algorithms
// It consists in iterate an array and switching the next value with the previously sorted values if needed
// Quadratic complexity O(n2) so it is not performative for high input volume
// Optimized if the input is partially sorted
// Can be combined with divide-to-conquer algorithms
func InsertionSort(A []int) []int {
	for i := 1; i < len(A); i++ {
		for y := i; y > 0; y-- {
			if A[y-1] > A[y] {
				A[y-1], A[y] = A[y], A[y-1]
				continue
			}
			break
		}
	}
	return A
}
