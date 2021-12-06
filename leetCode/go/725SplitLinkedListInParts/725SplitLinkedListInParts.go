package splitlinkedlistinparts

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
func splitListToParts(head *ListNode, k int) []*ListNode {
	index := [1000]*ListNode{}
	counter := 0
	for trace := head; trace != nil; trace = trace.Next {
		index[counter] = trace
		counter++
	}
	d := counter / k
	r := counter % k
	result := []*ListNode{}
	for i := 0; i < counter; {
		result = append(result, index[i])
		i = i + d
		if r > 0 {
			i += 1
			r--
		}
		index[i-1].Next = nil
	}
	for len(result) != k {
		result = append(result, nil)
	}
	return result
}
