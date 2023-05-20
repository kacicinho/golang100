type ListNode struct {
    Val  int
    Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    
    // Step 1: Find the middle point
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    // Step 2: Reverse the second half
    secondHalf := reverseList(slow)
    
    // Step 3: Compare the first half and the reversed second half
    firstHalf := head
    for firstHalf != nil && secondHalf != nil {
        if firstHalf.Val != secondHalf.Val {
            return false
        }
        firstHalf = firstHalf.Next
        secondHalf = secondHalf.Next
    }
    
    return true
}

// Helper function to reverse a linked list
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    current := head
    
    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    
