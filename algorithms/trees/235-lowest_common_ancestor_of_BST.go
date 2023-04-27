package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC: O(logn) - height of the tree
// SC: O(logn) - Create a call stack that potentially could be the height of the tree
func lowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
	// Check if both children are actually left of the root
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	// Check if both children are right of the root
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	// If children are split left and right; found ancestor
	return root
}

// 2023-04-26
// TC: O(logn) - Height of the tree
// SC: O(1) - Pointer arithmetic
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root

	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}

	return cur
}

func main() {
}
