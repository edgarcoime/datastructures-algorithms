package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2023-04-26
// TC: O(n) - Iterating through the list only once
// SC: O(1) - No extra datastructures needed just pointer logic
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	s := head
	f := head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

func main() {
}
