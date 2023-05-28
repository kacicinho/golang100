/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumOfLeftLeaves(root *TreeNode) int {
    stack := []*TreeNode{root}
    sum := 0

    for len(stack) > 0 {
        curr := stack[0]
        stack = stack[1:]

        if curr.Left != nil {
            if curr.Left.Left == nil && curr.Left.Right == nil {
                sum += curr.Left.Val
            } else {
                stack = append(stack, curr.Left)
            }
        }

        if curr.Right != nil {
            stack = append(stack, curr.Right)
        }
    }

    return sum
}
