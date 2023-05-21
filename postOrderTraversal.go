/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func postorderTraversal(root *TreeNode) []int {
    list := []int{} 

    if root == nil {
        return nil
    }

    left := postorderTraversal(root.Left)
    list = append(list, left...)

    right := postorderTraversal(root.Right)
    list = append(list, right...)

    list = append(list, root.Val)

    return list 
}