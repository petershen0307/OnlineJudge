package binarytreeinordertraversal

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
	the main idea is that visit left node first, when there isn't any left node then visit right
	when the right node is null then go to parent node

	0. if root is null, end
	1. if left node is not null, put current node to stack and go to left node
	1.1 if left node is null, record current node to result then go to right node (we record the current node to result, there is no need to put current node to stack)
	2. if right node is not null, go to right node and go back to step 1 to check the left node (left node should traverse first )
	2.1 if right node is null, go the parent node
	3. pop the node from stack and go to step 2 to check right node (we don't go to check left because we already traversed the left node)
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	inorder := make([]int, 0)
	index := root
	// check input is not nil root
	if nil == index {
		return inorder
	}
	type action int
	const (
		popNode action = iota
		goLeft
		goRight
	)
	theAction := goLeft
	// the stack is empty and the action is pop node from stack
	// this situation mean the tree traverse done
	for len(stack) != 0 || theAction != popNode {
		switch theAction {
		case popNode:
			theAction = goRight
			index = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inorder = append(inorder, index.Val)
		case goLeft:
			if index.Left != nil {
				stack = append(stack, index)
				index = index.Left
			} else {
				inorder = append(inorder, index.Val)
				theAction = goRight
			}
		case goRight:
			if index.Right != nil {
				index = index.Right
				theAction = goLeft
			} else {
				theAction = popNode
			}
		}
	}
	return inorder
}

func inorderTraversal2(root *TreeNode) []int {
	stack := list.New()
	traceResult := make([]int, 0)
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		popNode := stack.Back()
		root = stack.Remove(popNode).(*TreeNode)
		traceResult = append(traceResult, root.Val)
		root = root.Right
	}
	return traceResult
}
