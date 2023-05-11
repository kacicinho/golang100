/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /* recursive way*/
 func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return symmetricHelper(root.Left, root.Right)
}

func symmetricHelper(left *TreeNode, right *TreeNode) bool{
        if (left == nil && right == nil){
            return true
        }

        if (left == nil || right == nil){
            return false
        }

        if (left.Val != right.Val){
            return false
        }

        return symmetricHelper(left.Left, right.Right) && symmetricHelper(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    queue := []*TreeNode{root.Left, root.Right}
    
    for len(queue) > 0 {
        leftNode := queue[0]
        queue = queue[1:]
        rightNode := queue[0]
        queue = queue[1:]
        
        if leftNode == nil && rightNode == nil {
            continue
        }
        
        if leftNode == nil || rightNode == nil {
            return false
        }
        
        if leftNode.Val != rightNode.Val {
            return false
        }
        
        queue = append(queue, leftNode.Left, rightNode.Right, leftNode.Right, rightNode.Left)
    }
    
    return true
}