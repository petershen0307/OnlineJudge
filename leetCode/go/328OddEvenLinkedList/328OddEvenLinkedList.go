package oddevenlinkedlist

import . "github.com/petershen0307/OnlineJudge/tree/master/leetCode/go/general"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var evenHead *ListNode
	var evenTail *ListNode
	var lastNode *ListNode
	// let head be odd head
	for oddTail := head; oddTail != nil; oddTail = oddTail.Next {
		if evenHead == nil {
			evenHead = oddTail.Next
			evenTail = evenHead
		} else {
			evenTail.Next = oddTail.Next
			evenTail = evenTail.Next
		}
		lastNode = oddTail
		if evenTail != nil {
			oddTail.Next = evenTail.Next
			evenTail.Next = nil
		}
	}
	lastNode.Next = evenHead
	return head
}
