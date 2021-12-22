package recorderlist

import . "github.com/petershen0307/OnlineJudge/tree/master/leetCode/go/general"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	nodeIndex := []*ListNode{}
	for run := head; run != nil; run = run.Next {
		nodeIndex = append(nodeIndex, run)
	}
	previous := (*ListNode)(nil)
	for i, j := 0, len(nodeIndex)-1; i <= j; {
		nodeIndex[i].Next = nodeIndex[j]
		if previous != nil {
			previous.Next = nodeIndex[i]
		}
		previous = nodeIndex[j]
		previous.Next = nil
		i++
		j--
	}
}
