package swappingnodesinalinkedlist

import (
	. "github.com/petershen0307/OnlineJudge/tree/master/leetCode/go/general"
)

func swapNodes(head *ListNode, k int) *ListNode {
	var beginK, endK *ListNode
	countDown := k
	// do the reverse and get begin k node
	for trace := head; trace != nil; countDown-- {
		// the list is 1-indexed
		if countDown == 1 {
			beginK = trace
		}
		nextNode := trace.Next
		if head == trace {
			trace.Next = nil
		} else {
			trace.Next = head
			head = trace
		}
		trace = nextNode
	}
	// do the reverse again and get end k node
	countDown = k
	for trace := head; trace != nil; countDown-- {
		// the list is 1-indexed
		if countDown == 1 {
			endK = trace
		}
		nextNode := trace.Next
		if head == trace {
			trace.Next = nil
		} else {
			trace.Next = head
			head = trace
		}
		trace = nextNode
	}
	beginK.Val, endK.Val = endK.Val, beginK.Val
	return head
}
