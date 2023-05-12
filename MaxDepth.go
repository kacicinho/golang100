func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
	
    if leftDepth > rightDepth {
        return leftDepth + 1
    } else {
        return rightDepth + 1
    }
}


func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var depth int
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[i]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        queue = queue[size:]
        depth++
    }
    return depth
}