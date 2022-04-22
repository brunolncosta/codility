package algoritms

// SelectionSort
// It consists in iterate thw whole array over and over
// searching for the next position number, when found it switch their positions.
// Quadratic complexity O(n2) so it is not performative for high input volume
// Not optimized for any data scenario
// Maximum position swap is N (length of the input)
func SelectionSort(A []int) []int {
	N := len(A)
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if A[j-1] > A[j] {
				min = j
			}
		}
		A[i], A[min] = A[min], A[i]
	}
	return A
}
