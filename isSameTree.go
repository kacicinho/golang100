/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

 parcourir en profondeur l'arbre principale 
 si jamais un noeud == racine de la cible 
 alors je commence a parcourir la cible 
 je vérifie que le parcours en profondeur de la cible == le parcours de la racine 
 si jamais c'est good on continue sinon on reset 
 */


 func isSameTree(first *TreeNode, second *TreeNode){
    queue := *TreeNode[]{first, second}

    for len(queue) > 0 {
        node1 := queue[len(queue)-1]
        node2 := queue[len(queue)-2]
        queue := queue[:len(queue)-2]

        if node1 == nil || node2 == nil {
            return 
        }

        if node1.Val != node2.Val {
            return false 
        }

        queue = append(queue, node1.Left, node2.Left, node1.Right, node2.Right)

    }
}


func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    

}