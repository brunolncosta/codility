package equilieader

func Solution(A []int) int {
	leader := -1
	count := make(map[int]int)
	for _, a := range A {
		count[a] = count[a] + 1
		if count[a] > len(A)/2 {
			leader = a
		}
	}

	right := len(A)
	left := 0
	leadersRight := count[leader]
	leadersLeft := 0
	equiLeaders := 0
	for _, a := range A {
		right--
		left++
		if a == leader {
			leadersRight--
			leadersLeft++
		}

		if right/2 < leadersRight && left/2 < leadersLeft {
			equiLeaders++
		}
	}
	return equiLeaders
}
