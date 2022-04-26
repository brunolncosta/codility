package stonewall

func Solution(H []int) int {
	if len(H) <= 0 {
		return 0
	}

	block := 1
	stack := make([]int, 0)
	for i := 1; i < len(H); i++ {
		if H[i] == H[i-1] {
			continue
		}

		if H[i] > H[i-1] {
			stack = append(stack, H[i-1])
			block++
			continue
		}

		useBase := false
		for len(stack) > 0 {
			base := stack[len(stack)-1]
			if base == H[i] {
				useBase = true
				break
			}

			if base < H[i] {
				break
			}
			stack = stack[:len(stack)-1]
		}

		if !useBase {
			block++
		}
	}
	return block
}
