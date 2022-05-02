package binarytree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Exists(t *testing.T) {
	tree := NewTree()
	tree.Add(100)
	tree.Add(99)
	tree.Add(150)
	tree.Add(1)

	assert.True(t, tree.Exists(100))
	assert.True(t, tree.Exists(99))
	assert.True(t, tree.Exists(150))
	assert.True(t, tree.Exists(1))
	assert.False(t, tree.Exists(102))
}

func Test_ExistsBreadth(t *testing.T) {
	tree := NewTree()
	tree.Add(100)
	tree.Add(99)
	tree.Add(150)
	tree.Add(1)

	assert.True(t, tree.ExistsBreadth(100))
	assert.True(t, tree.ExistsBreadth(99))
	assert.True(t, tree.ExistsBreadth(150))
	assert.True(t, tree.ExistsBreadth(1))
	assert.False(t, tree.ExistsBreadth(102))
}

//             0
//      -6          100
//          -3   80     110
//                           120
func Test_GetLeftView(t *testing.T) {
	tree := NewTree()
	tree.Add(0)
	tree.Add(-6)
	tree.Add(-3)
	tree.Add(100)
	tree.Add(80)
	tree.Add(110)
	tree.Add(120)

	assert.True(t, tree.Exists(0))
	assert.True(t, tree.Exists(-6))
	assert.True(t, tree.Exists(-3))
	assert.True(t, tree.Exists(100))
	assert.True(t, tree.Exists(80))
	assert.True(t, tree.Exists(110))
	assert.True(t, tree.Exists(120))

	leftView := tree.GetLeftView()
	expected := []int{0, -6, -3, 120}
	assert.Equal(t, expected, leftView)
}

//                0
//        -6            100
//            -3     80     110
//                              120
func Test_GetLeftViewBreadth(t *testing.T) {
	tree := NewTree()
	tree.Add(0)
	tree.Add(-6)
	tree.Add(-3)
	tree.Add(100)
	tree.Add(80)
	tree.Add(110)
	tree.Add(120)

	assert.True(t, tree.ExistsBreadth(0))
	assert.True(t, tree.ExistsBreadth(-6))
	assert.True(t, tree.ExistsBreadth(-3))
	assert.True(t, tree.ExistsBreadth(100))
	assert.True(t, tree.ExistsBreadth(80))
	assert.True(t, tree.ExistsBreadth(110))
	assert.True(t, tree.ExistsBreadth(120))

	leftView := tree.GetLeftViewBreadth()
	expected := []int{0, -6, -3, 120}
	assert.Equal(t, expected, leftView)
}

//                  0
//           100                -6
//      110        80       -3
//  120
func Test_Revert(t *testing.T) {
	tree := NewTree()
	tree.Add(0)
	tree.Add(-6)
	tree.Add(-3)
	tree.Add(100)
	tree.Add(80)
	tree.Add(110)
	tree.Add(120)

	tree.Revert()
	leftView := tree.GetLeftViewBreadth()
	expected := []int{0, 100, 110, 120}
	assert.Equal(t, expected, leftView)
}

//                0
//        -6            100
//            -3     80     110
//                              120
func Test_GetButtonView(t *testing.T) {
	tree := NewTree()
	tree.Add(0)
	tree.Add(-6)
	tree.Add(-3)
	tree.Add(100)
	tree.Add(80)
	tree.Add(110)
	tree.Add(120)

	leftView := tree.GetButtonView()
	expected := []int{-6, -3, 0, 80, 100, 110, 120}
	assert.Equal(t, expected, leftView)
}
