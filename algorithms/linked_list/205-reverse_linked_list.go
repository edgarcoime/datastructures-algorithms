package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLL(l []int) *ListNode {
	head := &ListNode{}
	curr := head

	for _, num := range l {
		node := &ListNode{Val: num}
		curr.Next = node
		curr = node
	}

	// Exclude dummy head
	return head.Next
}

func printLL(l *ListNode) {
	if l == nil {
		// Hit at the end
		fmt.Printf("nil\n")
		fmt.Printf("nil")
		return
	}

	// Traverse Stack In Order
	fmt.Printf("%d -> ", l.Val)

	printLL(l.Next)

	// Traverse Stack in Reverse (Bubble back up)
	fmt.Printf(" <- %d", l.Val)
	return
}

// TC: O(n) - iterating through linked list
// SC: O(1) - Use of pointers w/o datastructures
func reverseList(head *ListNode) *ListNode {
	// Keep track of following node
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next

		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// TC: O(n) - Iterating through linked list
// SC: O(n) - Extra memory since stack frames
func reverseListRecursive(head *ListNode) *ListNode {
	// Base Case
	if head == nil || head.Next == nil {
		return head
	}

	// Divide into smaller so into the Base Case
	reversedListHead := reverseListRecursive(head.Next)

	// Operate on current stack frame
	head.Next.Next = head
	head.Next = nil
	return reversedListHead
}

// 1 ->
// 1 -> 2
// 1 -> 2 -> 3 ->

func main() {
	h := createLL([]int{1, 2, 3})
	// h := createLL([]int{1, 2, 3, 4, 5})
	// h := createLL([]int{1, 2, 3, 4, 5})
	// println(reverseListRecursive(h))
	printLL(h)
}
