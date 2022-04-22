package distinct

import "sort"

// SolutionWithSet The optimized solution does use SET instead of sort
func SolutionWithSet(A []int) int {
	set := make(map[int]bool)
	for _, a := range A {
		set[a] = true
	}

	return len(set)
}

// SolutionWithSort Golang sort implement O(nlogn)
func SolutionWithSort(A []int) int {
	sort.Ints(A)
	count := 0
	if len(A) > 0 {
		count = 1
	}

	for i := 1; i < len(A); i++ {
		if A[i-1] != A[i] {
			count++
		}
	}
	return count
}

// SolutionWithInsertionSort If you could not use goland sorting
// You can use insertions sort. It is slow O(n2) but will do the job.
func SolutionWithInsertionSort(A []int) int {
	for i := 1; i < len(A); i++ {
		for j := i - 1; j >= 0; j-- {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
	}

	count := 0
	if len(A) > 0 {
		count = 1
	}

	for i := 1; i < len(A); i++ {
		if A[i-1] != A[i] {
			count++
		}
	}
	return count
}
