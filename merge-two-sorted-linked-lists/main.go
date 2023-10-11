/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head:= &ListNode{}
    cur:= head

    for list1!=nil || list2!=nil {
        if list1 == nil {
            // append rest of list2 to curr and set that to nil to break
            cur.Next = list2
            list2 = nil
            continue
        }
        if list2 == nil {
            cur.Next = list1
            list1 = nil
            continue
        }

        if list1.Val > list2.Val {
            cur.Next = list2
            cur = cur.Next
            list2 = list2.Next
        } else {
            cur.Next = list1
            cur = cur.Next
            list1 = list1.Next
        }
    }
    return head.Next
}