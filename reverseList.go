func reverseList(head *ListNode) *ListNode {
    var stack []*ListNode

    // Push each node of the linked list onto the stack
    for head != nil {
        stack = append(stack, head)
        head = head.Next
    }

    // Update the Next pointers of the nodes in reverse order
    for i := len(stack) - 1; i > 0; i-- {
        stack[i].Next = stack[i-1]
    }

    // Set the Next of the original head to nil
    if len(stack) > 0 {
        stack[0].Next = nil
    }

    // Return the new head of the reversed list
    if len(stack) > 0 {
        return stack[len(stack)-1]
    }

    return nil
}
