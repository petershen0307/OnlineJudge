package general

import (
	"runtime"
)

func GetCallerName() string {
	// the grandparent is the actual caller we interest
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return funcName
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToLinkedList(input []int) *ListNode {
	var head *ListNode
	index := &head
	for _, v := range input {
		if *index == nil {
			*index = &ListNode{Val: v, Next: nil}
		}
		index = &(*index).Next
	}
	return head
}

func LinkedListToSlice(head *ListNode) []int {
	output := []int{}
	for ; head != nil; head = head.Next {
		output = append(output, head.Val)
	}
	return output
}
