package q1171

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

//https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/discuss/366350/C%2B%2B-O(n)-(explained-with-pictures)
func removeZeroSumSublists(head *ListNode) *ListNode {
	// add a root node before head to fix head will be removed
	root := &ListNode{Val: 0, Next: head}
	// store accumulative result
	aMap := make(map[int]*ListNode, 0)
	accumulateR := 0
	aMap[0] = root
	for head != nil {
		accumulateR += head.Val
		if v, found := aMap[accumulateR]; found {
			fromMap := v
			accumulateInMap := accumulateR
			for fromMap != head {
				// should delete from map next reference
				fromMap = fromMap.Next
				accumulateInMap += fromMap.Val
				// delete no longer reference in map
				// don't delete this val(accumulateR) in map
				if fromMap != head {
					delete(aMap, accumulateInMap)
				}
			}
			v.Next = head.Next
		} else {
			aMap[accumulateR] = head
		}
		head = head.Next
	}
	return root.Next
}
