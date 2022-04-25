package fish

func Solution(A []int, B []int) int {
	N := len(A)
	if N <= 0 {
		return 0
	}

	stack := make([]int, 0)
	sizes := append(stack, A[0])
	dirs := append(stack, B[0])

	for i := 1; i < N; i++ {
		size := A[i]
		dir := B[i]
		currEaten := false

		for len(sizes) > 0 {
			lastSize := sizes[len(sizes)-1]
			lastDir := dirs[len(dirs)-1]

			// 0 1 stack
			// 1 1 stack
			// 1 0 eat
			// 0 0 stack
			if dir != lastDir && lastDir != 0 {
				if size > lastSize {
					sizes = sizes[:len(sizes)-1]
					dirs = dirs[:len(dirs)-1]
					continue // eat every fish

				} else if size < lastSize {
					currEaten = true
				}
			}
			break
		}

		if !currEaten {
			dirs = append(dirs, dir)
			sizes = append(sizes, size)
		}
	}
	return len(sizes)
}
