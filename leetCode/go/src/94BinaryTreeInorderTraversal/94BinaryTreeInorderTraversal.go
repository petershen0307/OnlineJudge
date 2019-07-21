/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	inorder := make([]int, 0)
	index := root
	if nil == index {
		return inorder
	}
	bPopNode := false
	for len(stack) != 0 || !bPopNode {
		if bPopNode {
			bPopNode = false
			index = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inorder = append(inorder, index.Val)
			if index.Right != nil {
				index = index.Right
			} else {
				bPopNode = true
			}
			continue
		}

		if index.Left != nil {
			stack = append(stack, index)
			index = index.Left
			continue
		} else {
			inorder = append(inorder, index.Val)
		}
		if index.Right != nil {
			index = index.Right
		} else {
			bPopNode = true
		}
	}
	return inorder
}