package rotatelist

import . "github.com/petershen0307/OnlineJudge/tree/master/leetCode/go/general"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	nodeCount := 0
	var lastNode *ListNode = nil
	for trace := head; trace != nil; trace = trace.Next {
		nodeCount++
		if trace.Next == nil {
			lastNode = trace
		}
	}
	if nodeCount == 0 {
		return head
	}
	rotateNum := k % nodeCount
	for trace := head; trace != nil && nodeCount > rotateNum; trace = trace.Next {
		if nodeCount == rotateNum+1 {
			lastNode.Next = head
			head = trace.Next
			trace.Next = nil
			break
		}
		nodeCount--
	}
	return head
}
