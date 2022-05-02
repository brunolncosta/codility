package binarytree

type Tree struct {
	root *Node
}

func NewTree() Tree {
	return Tree{
		root: &Node{},
	}
}

func (tree Tree) Add(value int) {
	tree.root.add(value)
}

func (tree Tree) Exists(value int) bool {
	return tree.root.getNode(value, 0) != nil
}

func (tree Tree) ExistsBreadth(value int) bool {
	return tree.getNodeBreadth(value) != nil
}

// Breath first search - make no sense for binary trees
func (tree Tree) getNodeBreadth(value int) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, tree.root)
	for len(queue) > 0 {
		node := queue[0]

		if node.value == value {
			return node
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}

		if node.left != nil {
			queue = append(queue, node.left)
		}

		queue = queue[1:]
	}
	return nil
}

func (tree Tree) GetLeftViewBreadth() []int {
	queue := make([]*Node, 0)
	queue = append(queue, tree.root)
	leftView := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]

		if len(leftView)-1 < node.level {
			leftView = append(leftView, node.value)
		}

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}

		queue = queue[1:]
	}
	return leftView
}

func (tree Tree) GetLeftView() []int {
	result := make([]int, 0)
	tree.root.getLeftView(&result)
	return result
}

func (tree *Tree) Revert() {
	tree.root.revert()
}

func (tree Tree) GetButtonView() []int {
	return tree.root.getButtonView()
}

type Node struct {
	level       int
	value       int
	left, right *Node
}

func (node *Node) add(value int) {
	if value > node.value {
		if node.right != nil {
			node.right.add(value)
		} else {
			node.right = &Node{value: value, level: node.level + 1}
		}
		return
	}

	if value < node.value {
		if node.left != nil {
			node.left.add(value)
		} else {
			node.left = &Node{value: value, level: node.level + 1}
		}
		return
	}
}

// Depth First Search
func (node Node) getNode(value int, level int) *Node {
	if value > node.value {
		if node.right == nil {
			return nil
		}
		return node.right.getNode(value, node.level+1)
	}

	if value < node.value {
		if node.left == nil {
			return nil
		}
		return node.left.getNode(value, node.level+1)
	}

	return &node
}

func (node Node) getLeftView(leftView *[]int) {
	if len(*leftView)-1 < node.level {
		*leftView = append(*leftView, node.value)
	}

	if node.left != nil {
		node.left.getLeftView(leftView)
	}

	if node.right != nil {
		node.right.getLeftView(leftView)
	}
}

func (node *Node) revert() {
	if node.left != nil {
		node.left.revert()
	}

	if node.right != nil {
		node.right.revert()
	}

	aux := node.right
	node.right = node.left
	node.left = aux
}

func (node *Node) getButtonView() []int {

	var left []int
	if node.left != nil {
		left = node.left.getButtonView()
	}

	var right []int
	if node.right != nil {
		right = node.right.getButtonView()
	}

	buttonView := make([]int, 0)
	buttonView = append(buttonView, left...)
	buttonView = append(buttonView, node.value)
	buttonView = append(buttonView, right...)
	return buttonView
}
