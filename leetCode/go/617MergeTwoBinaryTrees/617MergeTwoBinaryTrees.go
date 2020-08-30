/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if nil == t1 {
        return t2
    }
    if nil == t2 {
        return t1
    }
    // add val to t1
    t1.Val += t2.Val
    bSkipRight := false
    bSkipLeft := false
    // copy t2 to t1
    if nil == t1.Right && nil != t2.Right{
        t1.Right = t2.Right
        bSkipRight = true
    }
    
    if nil == t1.Left && nil != t2.Left{
        t1.Left = t2.Left
        bSkipLeft = true
    }
    
    if !bSkipRight {
        mergeTrees(t1.Right, t2.Right)
    }
    if !bSkipLeft {
        mergeTrees(t1.Left, t2.Left)
    }
    return t1
}
