package countnondivisible

import "sort"

func Solution(A []int) []int {
	numbers := make(map[int]int)
	max := 0
	for _, a := range A {
		numbers[a] = 0
		if a > max {
			max = a
		}
	}

	countOne := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 1 {
			countOne++
			continue
		}

		for j := A[i]; j <= max; j += A[i] {
			if _, ok := numbers[j]; ok {
				numbers[j] += 1
			}
		}
	}

	result := make([]int, len(A))
	for i, n := range A {
		divisor := numbers[n]
		if countOne > 0 {
			divisor += countOne
		}

		result[i] = len(A) - divisor
	}
	return result
}

func SlowSolution(A []int) []int {

	var s []int
	s = append(s, A...)
	sort.Ints(s) // 1 2 3 3 6

	numbers := make(map[int]int)

	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && s[i] == s[i+1] {
			continue
		}
		count := 0
		for j := 0; j < i; j++ {
			if s[i]%s[j] == 0 {
				count++
			}
		}

		count = i - count
		nonDiv := count + (len(s) - (i + 1))
		numbers[s[i]] = nonDiv
	}

	result := make([]int, len(A))

	for i, n := range A {
		result[i] = numbers[n]
	}
	return result
}
