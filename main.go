package main

import (
	"fmt"
)

type TreeNode struct {
	Value  int
	Left   *TreeNode
	Right  *TreeNode
	Height int
}

func (t *TreeNode) insert(value int) {
	if t == nil {
		return
	} else if value <= t.Value {
		if t.Left == nil {
			t.Left = &TreeNode{Value: value}
		} else {
			t.Left.insert(value)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{Value: value}
		} else {
			t.Right.insert(value)
		}
	}
}

func (t *TreeNode) search(value int) *TreeNode {
	if t == nil || t.Value == value {
		return t
	} else if value < t.Value {
		return t.Left.search(value)
	} else {
		return t.Right.search(value)
	}
}

func (t *TreeNode) delete(value int) *TreeNode {
	if t == nil {
		return t
	}

	if value < t.Value {
		t.Left = t.Left.delete(value)
	} else if value > t.Value {
		t.Right = t.Right.delete(value)
	} else {
		if t.Left == nil {
			return t.Right
		} else if t.Right == nil {
			return t.Left
		}

		t.Value = t.Right.min().Value
		t.Right = t.Right.delete(t.Value)
	}

	return t
}

func (t *TreeNode) inOrderTraversal() {
	if t == nil {
		return
	}
	t.Left.inOrderTraversal()
	println(t.Value)
	t.Right.inOrderTraversal()
}

func (t *TreeNode) preOrderTraversal() {
	if t == nil {
		return
	}

	println(t.Value)
	t.Left.preOrderTraversal()
	t.Right.preOrderTraversal()
}

func (t *TreeNode) postOrderTraversal() {
	if t == nil {
		return
	}

	t.Left.postOrderTraversal()
	t.Right.postOrderTraversal()
	println(t.Value)
}

func (t *TreeNode) balance() *TreeNode {
	if t == nil {
		return t
	}

	balance := t.getBalance()
	if balance > 1 {
		if t.Left.getBalance() < 0 {
			t.Left = t.Left.rotateLeft()
		}
		return t.rotateRight()
	} else if balance < -1 {
		if t.Right.getBalance() > 0 {
			t.Right = t.Right.rotateRight()
		}
		return t.rotateLeft()
	}

	return t
}

func (t *TreeNode) rotateLeft() *TreeNode {
	newRoot := t.Right
	t.Right = newRoot.Left
	newRoot.Left = t

	t.Height = 1 + max(height(t.Left), height(t.Right))
	newRoot.Height = 1 + max(height(newRoot.Left), height(newRoot.Right))

	return newRoot
}

func (t *TreeNode) rotateRight() *TreeNode {
	newRoot := t.Left
	t.Left = newRoot.Right
	newRoot.Right = t

	t.Height = 1 + max(height(t.Left), height(t.Right))
	newRoot.Height = 1 + max(height(newRoot.Left), height(newRoot.Right))

	return newRoot
}

func (t *TreeNode) getBalance() int {
	if t == nil {
		return 0
	}

	return height(t.Left) - height(t.Right)
}

func (t *TreeNode) min() *TreeNode {
	current := t
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(t *TreeNode) int {
	if t == nil {
		return 0
	}

	return t.Height
	//return 1 + max(height(t.Left), height(t.Right))
}

func main() {
	root := &TreeNode{Value: 10}
	values := []int{5, 15, 3, 7, 13, 17, 1, 9, 11, 14, 19}
	for _, value := range values {
		root.insert(value)
	}

	fmt.Println("In Order Traversal")
	root.inOrderTraversal()
	fmt.Println("Pre Order Traversal")
	root.preOrderTraversal()
	fmt.Println("Post Order Traversal")
	root.postOrderTraversal()

	searchVal := 14
	result := root.search(searchVal)
	if result != nil {
		fmt.Printf("Found %d\n", searchVal)
	} else {
		fmt.Printf("Not Found %d\n", searchVal)
	}

	deleteVal := 13
	root = root.delete(deleteVal)
	fmt.Printf("Deleted %d\n", deleteVal)

	fmt.Println("In Order Traversal after deletion.")
	root.inOrderTraversal()

}
