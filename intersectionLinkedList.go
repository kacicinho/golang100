/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func isPresent(target *ListNode, diffNodes []*ListNode) bool{
    for _, val := range diffNodes {
        if val == target {
            return true
        }
    }
    return false
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    diffNodes := []*ListNode{headA}

    for headA != nil {
        headA = headA.Next
        diffNodes = append(diffNodes, headA)
    }

    for headB != nil {
        headB = headB.Next
        if isPresent(headB, diffNodes) {
            return headB
        }
    }
    return nil
}