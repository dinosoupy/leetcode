package main

import (
	"fmt"
)

type ListNode struct {
  Val int
  Next *ListNode
}

type List struct {
	Head *ListNode
}

func main() {
	arr1 := [3]int{3,4,2}
	arr2 := [3]int{7,5,6}

	l1 := List{}
	for i := 0; i < 3; i++ {
		l1.Insert(arr1[i])
	}
	fmt.Printf("This is the first List: ")
	Show(&l1)

	l2 := List{}
	for i := 0; i < 3; i++ {
		l2.Insert(arr2[i])
	}
	fmt.Printf("This is the second List: ")
	Show(&l2)

	resultHead := addTwoNumbers(l1.Head,l2.Head)
	fmt.Printf("This is the resulting List: ")
	for resultHead.Next!=nil {
		fmt.Print(resultHead.Val)
		resultHead = resultHead.Next
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Make a new ListNode and store its first element in head
	head := new(ListNode)
	// Make an interator element for traversing ListNodes (at first, it's head)
	iter := head
	carry := 0
	// Start iterating if either l1 or l2 is non-empty or if there's a carry
	if l1 != nil || l2 != nil || carry != 0 {
		// Assign 0 to nil nodes
		if l1 == nil {
			l1.Val = 0
		}
		if l2 == nil {
			l2.Val = 0
		}

		fmt.Printf("l1 = \nl2 = \n", l1.Val, l2.Val)
		// Add values and determine carry
		sum := carry + l1.Val + l2.Val
		iter.Next = &ListNode{Val: sum % 10, Next: nil}
		carry = sum / 10

		// Advance l1, l2
		l1 = l1.Next
		l2 = l2.Next
	}
	return head
}

func (l *List) Insert(v int) {
	list := &ListNode{Val: v, Next: nil}
	if l.Head == nil {
		l.Head = list
	} else {
		p := l.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = list
	}
}

func Show(l *List) {
	p := l.Head
	for p != nil {
		fmt.Printf("-> %v ", p.Val)
		p = p.Next
	}
}
