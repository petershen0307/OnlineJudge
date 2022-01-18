package recoverbinarysearchtree

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
func recoverTree(root *TreeNode) {
	array := []*TreeNode{}
	inOrderTraverse(root, &array)
	first := -1
	second := -1

	for prev, i := 0, 1; i != len(array); i++ {
		// 1 3 2 4
		// first, second 會連續被改值
		if array[prev].Val > array[i].Val && first == -1 {
			first = prev
		}
		// 1 4 3 2 5 6
		if array[prev].Val > array[i].Val {
			second = i
		}
		prev = i
	}
	array[first].Val, array[second].Val = array[second].Val, array[first].Val
}

func inOrderTraverse(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	inOrderTraverse(root.Left, nodes)
	(*nodes) = append((*nodes), root)
	inOrderTraverse(root.Right, nodes)
}
