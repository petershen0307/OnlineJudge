package maximumbinarytreeii

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

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	// s1. bigger than root, the val be the root, all subtree to left node, cause val is the last value
	// s2. go right child to find val will be the root of subtree, insert subtree to left of val
	if nil == root {
		return nil
	}

	newRoot := new(TreeNode)
	if val == 0 {
		// copy node
		newRoot.Val = root.Val
		newRoot.Left = insertIntoMaxTree(root.Left, 0)
		newRoot.Right = insertIntoMaxTree(root.Right, 0)
		return newRoot
	}

	if val > root.Val {
		newRoot.Val = val
		newRoot.Right = nil
		newRoot.Left = insertIntoMaxTree(root, 0)
		return newRoot
	}
	newRoot.Val = root.Val
	newRoot.Left = insertIntoMaxTree(root.Left, 0)
	newRoot.Right = insertIntoMaxTree(root.Right, val)
	// append val to right
	if nil == newRoot.Right {
		newRoot.Right = new(TreeNode)
		newRoot.Right.Val = val
		newRoot.Right.Left = nil
		newRoot.Right.Right = nil
	}
	return newRoot
}
