package binarytreerightsideview

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

func rightSideView(root *TreeNode) []int {
	// use level order traversal
	rightSideViewValues := []int{}
	nodes := [][]*TreeNode{}
	if root != nil {
		nodes = append(nodes, []*TreeNode{root})
	}
	for i := 0; i != len(nodes); i++ {
		temp := []*TreeNode{}
		rightSideViewValues = append(rightSideViewValues, nodes[i][0].Val)
		for _, node := range nodes[i] {
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
		}
		if len(temp) > 0 {
			nodes = append(nodes, temp)
		}
	}
	return rightSideViewValues
}
