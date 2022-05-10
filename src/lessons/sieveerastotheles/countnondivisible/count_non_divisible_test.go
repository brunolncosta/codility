package countnondivisible

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Solution(t *testing.T) {
	assert.Equal(t, []int{2, 4, 3, 2, 0}, Solution([]int{3, 1, 2, 3, 6}))
	assert.Equal(t, []int{1, 1}, Solution([]int{3, 4}))
	assert.Equal(t, []int{0, 1}, Solution([]int{4, 2}))
	assert.Equal(t, []int{0, 0}, Solution([]int{100, 100}))
	assert.Equal(t, []int{0, 0, 0, 0}, Solution([]int{100, 100, 100, 100}))
	assert.Equal(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Solution([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
	assert.Equal(t, []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}, Solution([]int{3, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
	assert.Equal(t, []int{4, 0, 0, 0, 0}, Solution([]int{1, 100, 100, 100, 100}))
}
