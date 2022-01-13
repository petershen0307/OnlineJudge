package insertintoabinarysearchtree

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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	insert(&root, val)
	return root
}

func insert(rootPtr **TreeNode, val int) {
	if (*rootPtr) == nil {
		(*rootPtr) = &TreeNode{Val: val}
		return
	}
	if val < (*rootPtr).Val {
		insert(&(*rootPtr).Left, val)
	} else {
		insert(&(*rootPtr).Right, val)
	}
}
