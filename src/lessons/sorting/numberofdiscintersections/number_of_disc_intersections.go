package numberofdiscintersections

import "sort"

// Solution Was very difficult to understand the problem.
// I had to transform the problem in a simple one using sorting,
// and then use stack to solve it
func Solution(A []int) int {
	if len(A) <= 1 {
		return 0
	}

	end := make([]int, len(A))
	start := make([]int, len(A))

	for i := range A {
		end[i] = i + A[i]
		start[i] = i - A[i]
	}

	sort.Ints(end)
	sort.Ints(start)

	parentheses := make([]rune, 0)
	i := 0
	j := 0
	for i < len(A) && j < len(A) {
		if start[i] > end[j] {
			parentheses = append(parentheses, ')')
			j++
		} else {
			parentheses = append(parentheses, '(')
			i++
		}
	}

	for ; i < len(A); i++ {
		parentheses = append(parentheses, '(')
	}

	for ; j < len(A); j++ {
		parentheses = append(parentheses, ')')
	}

	stack := make([]rune, 0)
	intersection := 0
	for _, p := range parentheses {
		if p == '(' {
			stack = append(stack, '(')
			intersection = intersection + (len(stack) - 1)
		} else {
			stack = stack[:len(stack)-1]
		}
		if intersection > 10000000 {
			return -1
		}
	}
	return intersection
}
