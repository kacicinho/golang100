/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSameTree(p *TreeNode, q *TreeNode) bool {
    pQueue := []*TreeNode{p}
    qQueue := []*TreeNode{q}
    for len(pQueue) > 0 && len(qQueue) > 0 {
        e1 := pQueue[0]
        e2 := qQueue[0]

        pQueue = pQueue[1:]
        qQueue = qQueue[1:]

        if e1 == nil && e2 == nil {
            continue
        }
        
        if e1 == nil || e2 == nil || e2.Val != e1.Val {
            return false
        }

        pQueue = append(pQueue, e1.Left, e1.Right)
        qQueue = append(qQueue, e2.Left, e2.Right)
    }

    return true 
}