package partitionlist

import (
	. "github.com/petershen0307/OnlineJudge/tree/master/leetCode/go/general"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var bigHead *ListNode
	var big **ListNode
	var small **ListNode
	var smallHead *ListNode
	for run := head; run != nil; run = run.Next {
		if run.Val >= x {
			if bigHead == nil {
				bigHead = run
			}
			if big != nil {
				(*big) = run
			}
			big = &run.Next
		} else {
			if smallHead == nil {
				smallHead = run
			}
			if small != nil {
				(*small) = run
			}
			small = &run.Next
		}
	}
	if big != nil {
		(*big) = nil
	}
	if small != nil {
		(*small) = bigHead
	} else {
		smallHead = bigHead
	}
	return smallHead
}
