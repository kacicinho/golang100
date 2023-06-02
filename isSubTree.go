/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    visitedRoot := []*TreeNode{root}
    possibleChildren := []*TreeNode{}
    maybeIsSubtree := false

    for len(visited)>0 {

        curr := visitedRoot[0]

        if curr.Left != nil {
            visitedRoot = append(visitedRoot, curr.Left)
        }

        if curr.Right != nil {
            visitedRoot = append(visitedRoot, curr.Right)
        }

        if curr.Value == subRoot.Value= {
            maybeIsSubtree = true 
        } else {
            maybeIsSubtree = false 
        }

        visitedRoot = visitedRoot[1:]

    }


}