package sorting

// BubbleSort Very simple sort algorithms
// It consists in reverse iterate the array and reserve position if needed.
// The swap will make the sorted number go to the up part of the array like bubbles rising.
// It has the same advantages and problem of insertion sorting.
func BubbleSort(A []int) []int {
	N := len(A)
	for i := N - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if A[j+1] < A[j] {
				A[j+1], A[j] = A[j], A[j+1]
			}
		}
	}
	return A
}
