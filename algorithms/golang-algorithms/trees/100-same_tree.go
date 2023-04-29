package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2023-04-27
// TC: O(p+q) - Length of p and q trees
// SC: O(p | q) - Whichever tree is the shortest
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Both are nil then same
	if p == nil && q == nil {
		return true
	}

	// If either are nil but not the other then not the same
	if p == nil || q == nil {
		return false
	}

	// If values are not the same then not same
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
}
