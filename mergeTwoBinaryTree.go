/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    track1 := []*TreeNode{root1, root2}

    


    for len(track1) > 0 {
        node1 := track1[0]
        node2 := track1[1]

        track1 = track1[2:]

        node1.Val += node2.Val

        if node1.Left != nil && node2.Left != nil {
            track1 = append(track1, node1.Left, node2.Left)
        } else if node1.Left == nil {
            node1.Left = node2.Left 
        }

        if node1.Right != nil && node2.Right != nil {
            track1 = append(track1, node1.Right, node2.Right)
        } else if node1.Right == nil {
            node1.Right = node2.Right 
        } 
    }

    return root1
}