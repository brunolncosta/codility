package nesting

func Solution(S string) int {
	lettersMap := map[rune]rune{
		'(': ')',
		')': '(',
	}

	letters := []rune(S)
	stack := make([]rune, 0)
	for _, l := range letters {
		if _, exists := lettersMap[l]; !exists {
			return 0
		}

		if l == ')' {
			if len(stack) <= 0 {
				return 0
			}
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, l)

	}
	if len(stack) == 0 {
		return 1
	}
	return 0
}
