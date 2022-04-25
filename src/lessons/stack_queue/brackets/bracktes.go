package brackets

func Solution(S string) int {
	if S == "" {
		return 1
	}

	closeMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	openMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := make([]rune, 0, len(S))
	letters := []rune(S)

	for _, l := range letters {
		if openLetter, exists := closeMap[l]; exists {
			if len(stack) <= 0 {
				return 0 // if closing with none openMap it is not nested
			}

			lastLetter := stack[len(stack)-1]
			if openLetter == lastLetter {
				stack = stack[:len(stack)-1] // pop
				continue
			}
			return 0 // itt a closeMap not nested
		}

		if _, exists := openMap[l]; !exists {
			return 0 // invalid character
		}
		stack = append(stack, l) // push
	}

	if len(stack) == 0 {
		return 1 // only return 1 if there is expression opened
	}

	return 0
}
