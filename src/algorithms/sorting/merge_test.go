package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "random",
			args: []int{10, 6, 15, 4, 1, 45},
			want: []int{1, 4, 6, 10, 15, 45},
		},
	}

	for _, test := range tests {
		sorted := MergeSort(test.args)
		assert.Equal(t, test.want, sorted)
	}
}
