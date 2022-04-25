package sorting

func MergeSort(A []int) []int {
	N := len(A)
	if N < 2 {
		return A
	}

	right := MergeSort(A[N/2:])
	left := MergeSort(A[:N/2])
	return merge(right, left)
}

func merge(right []int, left []int) []int {
	merged := make([]int, 0, len(right)+len(left))
	r := 0
	l := 0
	for r < len(right) && l < len(left) {
		if right[r] < left[l] {
			merged = append(merged, right[r])
			r++
		} else {
			merged = append(merged, left[l])
			l++
		}
	}

	// It will last values for one array only. Or right or left will remain, never both.
	for ; r < len(right); r++ {
		merged = append(merged, right[r])
	}

	for ; l < len(left); l++ {
		merged = append(merged, left[l])
	}

	return merged
}
