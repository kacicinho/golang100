/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil 
    }
    modified := []*TreeNode{root}
    ret := root

    for len(modified) > 0 {

        curr := modified[0]
        modified = modified[1:]
        temp := curr.Left
        curr.Left = curr.Right
        curr.Right = temp

        if (curr.Left != nil) {
            modified = append(modified, curr.Left)
        } 
        if (curr.Right != nil) {
            modified = append(modified, curr.Right)
        }
        
    }

    return ret

}
