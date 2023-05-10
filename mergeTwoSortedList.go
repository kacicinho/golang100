/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
 The time complexity of the solution I provided for merging two sorted linked lists is already optimal at O(m+n), where m and n are the lengths of the input lists. This is because we have to visit each node in each list at least once in order to construct the merged list, and doing so requires O(m+n) time.

That being said, there are other algorithms for merging two sorted arrays or lists that have the same time complexity but might be faster in practice for certain types of data. One such algorithm is called merge sort, which is a divide-and-conquer algorithm that can be used to merge two sorted arrays or lists in O(m+n) time.

However, for small input sizes or in situations where the input lists are already partially sorted or contain duplicates, the simple iterative approach that we used in our solution might be faster in practice than more complex algorithms like merge sort. So, it's important to consider the specific characteristics of your input data when choosing an algorithm for merging two sorted lists.
*/

 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    list3 := &ListNode{}
    node := list3

    for (list1!= nil && list2!= nil) {

        if (list1.Val < list2.Val){
            node.Next = list1
            list1 = list1.Next
        } else{
            node.Next = list2
            list2 = list2.Next
        }

        node = node.Next
    }

    if list1 != nil {
        node.Next = list1
    } else if list2 != nil {
        node.Next = list2
    }

    return list3.Next
}