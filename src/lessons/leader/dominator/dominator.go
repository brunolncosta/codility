package dominator

func Solution(A []int) int {
	if len(A) <= 0 {
		return -1
	}

	count := make(map[int]int)
	half := len(A) / 2
	for i, item := range A {
		count[item] = count[item] + 1
		if count[item] > half {
			return i
		}
	}
	return -1
}
