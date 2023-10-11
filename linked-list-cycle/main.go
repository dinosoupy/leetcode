/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// SOLUTION 1: 
// Time: O(n)
// Space: O(n)
func hasCycle(head *ListNode) bool {
    visited:=make(map[*ListNode]struct{})

    if head==nil {return false}

    for head.Next!=nil {
        if _, exists:=visited[head.Next]; exists {
            return true
        }
        visited[head.Next] = struct{}{}
        head = head.Next
    }
    return false
}

// SOLUTION 2
// Time: O(n)
// Space: O(1)
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    // Time: O(n)
    for slow, fast := head, head; fast != nil && fast.Next != nil; {
        fast, slow = fast.Next.Next, slow.Next
        if fast == slow { return true }
    }
    return false
}