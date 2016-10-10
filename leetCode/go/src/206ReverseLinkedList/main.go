package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListIterator(head *ListNode) *ListNode {
	var reverseHead *ListNode
	var tmp *ListNode
	for i := head; i != nil; i = tmp {
		tmp = i.Next
		i.Next = reverseHead
		reverseHead = i
	}
	return reverseHead
}

func reverseListRecursive(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	reverseHead := reverseListRecursive(head.Next)
	var tail *ListNode
	for i := reverseHead; nil != i; i = i.Next {
		tail = i
	}
	tail.Next = head
	head.Next = nil
	return reverseHead
}

func print(head *ListNode) {
	for i := head; nil != i; i = i.Next {
		fmt.Print(*i, " ")
	}
	fmt.Println()
}

func main() {
	var head *ListNode
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	print(head)
	head = reverseListIterator(head)
	print(head)
	head = reverseListRecursive(head)
	print(head)
}
