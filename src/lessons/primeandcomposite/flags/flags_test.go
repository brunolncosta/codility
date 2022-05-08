package flags

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Solution(t *testing.T) {
	assert.Equal(t, 3, Solution([]int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}))
	assert.Equal(t, 0, Solution([]int{}))
	assert.Equal(t, 0, Solution([]int{5, 2, 1}))
	assert.Equal(t, 1, Solution([]int{1, 2, 1}))
	assert.Equal(t, 0, Solution([]int{1, 2, 3}))
	assert.Equal(t, 1, Solution([]int{1, 2, 3, 2}))
	assert.Equal(t, 1, Solution([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 3}))
	assert.Equal(t, 2, Solution([]int{0, 0, 0, 0, 0, 1, 0, 1, 0, 1}))
	assert.Equal(t, 4, Solution([]int{0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1}))
}
