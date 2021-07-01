package MergeTwoSortedLists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultHead *ListNode
	var resultIndex *ListNode
	finish := false
	for !finish {
		var temp *ListNode
		if l1 != nil && l2 != nil {
			if l1.Val > l2.Val {
				temp = l2
				l2 = l2.Next
			} else {
				temp = l1
				l1 = l1.Next
			}
		} else if l1 == nil {
			temp = l2
			finish = true
		} else {
			temp = l1
			finish = true
		}
		if resultHead == nil {
			resultHead = temp
			resultIndex = resultHead
		} else {
			resultIndex.Next = temp
			resultIndex = temp
		}
	}
	return resultHead
}
