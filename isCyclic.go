/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func isPresent(elt *ListNode, list []*ListNode) bool{
    for _, val := range list{
        if elt == val {
            return true
        }
    }
    return false
}

func hasCycle(head *ListNode) bool {
    curr := []*ListNode{head}
    if head == nil || head.Next == nil  {
        return false
    }
    head = head.Next
    for head != nil {
        if (isPresent(head, curr)) {
            return true
        } else {
            curr = append(curr, head)   
            head = head.Next
        }
        
    }
    return false
}