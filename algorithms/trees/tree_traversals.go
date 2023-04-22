package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (node *TreeNode) Insert(data int) *TreeNode {
	if node == nil {
		return New(data)
	}

	if data <= node.Val {
		node.Left = node.Left.Insert(data)
	} else {
		node.Right = node.Right.Insert(data)
	}

	return node
}

// Uses:
// - Values of nodes in sorted order
// - Reverse order can be done by visiting right first then left
func inOrder(n *TreeNode) {
	if n != nil {
		inOrder(n.Left)
		fmt.Printf("%d ", n.Val)
		inOrder(n.Right)
	}
}

// Uses:
// - Used to create a copy of the tree
// - Useful to get prefix expression of expression tree
func preOrder(n *TreeNode) {
	if n != nil {
		fmt.Printf("%d ", n.Val)
		preOrder(n.Left)
		preOrder(n.Right)
	}
}

// Uses:
// - Used to delete a tree
// - Useful to get postfix expression of expression tree
func postOrder(n *TreeNode) {
	if n != nil {
		preOrder(n.Left)
		preOrder(n.Right)
		fmt.Printf("%d ", n.Val)
	}
}

func main() {
	//         5
	//       /   \
	//     3       8
	//   /   \   /   \
	// 1     4  7     12
	node := New(5)
	node.Insert(3)
	node.Insert(8)
	node.Insert(1)
	node.Insert(4)
	node.Insert(7)
	node.Insert(12)
	println("InOrder:")
	inOrder(node)

	println("\nPreOrder:")
	preOrder(node)

	println("\nPostOrder:")
	postOrder(node)
}
