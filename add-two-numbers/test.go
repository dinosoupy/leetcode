// Reverse a linked list

type Node struct {
	Val int
	Next *Node
}

head := &Node{} // global variable

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}