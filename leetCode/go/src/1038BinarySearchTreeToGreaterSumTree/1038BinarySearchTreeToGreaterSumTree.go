/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
    total := 0
    traceBst(root, &total)
    return root
}

func traceBst(root *TreeNode, total *int) {
    if nil == root {
        return
    }
    traceBst(root.Right, total)
    ori := root.Val
    root.Val += *total
    *total += ori
    traceBst(root.Left, total)
}
