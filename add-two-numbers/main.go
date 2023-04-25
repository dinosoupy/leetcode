package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr1 := []int{9,9}
	arr2 := []int{9}

	l1 := toListNode(arr1)
	l2 := toListNode(arr2)

	resultHead := addTwoNumbers(l1, l2)
	fmt.Print("This is the resulting List: ")
	for resultHead != nil {
		fmt.Print(resultHead.Val)
		resultHead = resultHead.Next
	}
	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode) // Make a new ListNode head for the output
	iter := dummyHead // Make an iterator element for traversing ListNodes (at first, it's head)

	carry := 0
	// Start iterating if either l1 or l2 is non-empty
	for l1 != nil || l2 != nil {
		// Assign 0 to nil nodes in case one of the ListNodes is longer
		if l1 == nil && l2 != nil {
			l1 = &ListNode{Val: 0}
		}
		if l2 == nil && l1 != nil {
			l2 = &ListNode{Val: 0}
		}

		// Add and set values; determine carry
		sum := l1.Val + l2.Val + carry
		iter.Next = &ListNode{Val: sum % 10}
		carry = sum / 10

		// Advance l1, l2, and iter
		l1 = l1.Next
		l2 = l2.Next
		iter = iter.Next
	}
	// If there is a trailing carry, add that to the end
	if carry > 0 {
		iter.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}

// Helper function to convert arrays into ListNode for testing
func toListNode(arr []int) *ListNode {
	// Make a dummy head and iterator
	dummyHead := new(ListNode)
	iter := dummyHead

	// Loop through array items and add them as nodes in ListNode
	for i := 0; i < len(arr); i++ {
		iter.Next = &ListNode{Val: arr[i]}
		iter = iter.Next
	}
	return dummyHead.Next
}

