package distributecoinsinbinarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// solution: https://leetcode.com/problems/distribute-coins-in-binary-tree/discuss/221939/C%2B%2B-with-picture-post-order-traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	move := 0
	postOrder(root, &move)
	return move
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func postOrder(root *TreeNode, move *int) int {
	if root == nil {
		return 0
	}
	// 回傳值是負數表示缺少多少錢幣
	coinLeft := postOrder(root.Left, move)
	coinRight := postOrder(root.Right, move)
	*move += abs(coinLeft) + abs(coinRight)
	// 減一是保留給當前的節點
	return coinLeft + coinRight + root.Val - 1
}
