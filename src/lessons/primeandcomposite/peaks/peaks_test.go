package peaks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Solution(t *testing.T) {
	assert.Equal(t, 3, Solution([]int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}))
}
