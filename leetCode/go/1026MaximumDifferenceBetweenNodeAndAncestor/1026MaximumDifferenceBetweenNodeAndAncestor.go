package maximumdifferencebetweennodeandancestor

import (
	. "github.com/petershen0307/OnlineJudge/tree/master/leetCode/go/general"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	return traversal(root, root.Val, root.Val)
}

func traversal(root *TreeNode, max, min int) int {
	if root == nil {
		return max - min
	}
	if root.Val > max {
		max = root.Val
	}
	if root.Val < min {
		min = root.Val
	}
	sumL := traversal(root.Left, max, min)
	sumR := traversal(root.Right, max, min)
	if sumL > sumR {
		return sumL
	}
	return sumR
}
