package RangeSumofBST

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	v := 0
	if root.Val >= L && root.Val <= R {
		v = root.Val
	}
	l := rangeSumBST(root.Left, L, R)
	r := rangeSumBST(root.Right, L, R)
	return v + r + l
}
