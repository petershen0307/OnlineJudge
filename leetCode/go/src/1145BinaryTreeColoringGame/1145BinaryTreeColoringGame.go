package binarytreecoloringgame

import "container/list"

// TreeNode defined
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	stack := list.New()
	chosenParentCount := 0
	chosenLeftCount := 0
	chosenRightCount := 0

	for root != nil || stack.Len() != 0 {
		for root != nil && root.Val != x {
			chosenParentCount++
			stack.PushBack(root)
			root = root.Left
		}
		if root != nil && root.Val == x {
			chosenLeftCount = countSubtree(root.Left)
			//chosenRightCount = countSubtree(root.Right)
			root = nil
			continue
		}
		popNode := stack.Back()
		root = stack.Remove(popNode).(*TreeNode)
		root = root.Right
	}
	chosenRightCount = n - (chosenParentCount + chosenLeftCount + 1)
	return (chosenParentCount > 1+chosenLeftCount+chosenRightCount) ||
		(chosenLeftCount > 1+chosenParentCount+chosenRightCount) ||
		(chosenRightCount > 1+chosenLeftCount+chosenParentCount)
}

func countSubtree(root *TreeNode) int {
	counter := 0
	stack := list.New()
	for root != nil || stack.Len() != 0 {
		for root != nil {
			counter++
			stack.PushBack(root)
			root = root.Left
		}
		popNode := stack.Back()
		root = stack.Remove(popNode).(*TreeNode)
		root = root.Right
	}
	return counter
}
