/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func preorderTraversal(root *TreeNode) []int {
    list := []int{}

    if root == nil {
        return []int{}
    }
    list = append(list, root.Val)


    left := preorderTraversal(root.Left)
    list = append(list, left...)

    right := preorderTraversal(root.Right)
    list = append(list, right...)
    

    return list 
}