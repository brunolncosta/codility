package prime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Solution(t *testing.T) {
	assert.Equal(t, 1, Solution(2))
	assert.Equal(t, 1, Solution(3))
	assert.Equal(t, 0, Solution(4))
	assert.Equal(t, 1, Solution(5))
	assert.Equal(t, 0, Solution(6))
	assert.Equal(t, 1, Solution(7))
	assert.Equal(t, 1, Solution(13))
	assert.Equal(t, 1, Solution(17))
	assert.Equal(t, 0, Solution(20))
}
