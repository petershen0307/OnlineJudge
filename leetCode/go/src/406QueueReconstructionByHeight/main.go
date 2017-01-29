package main

import (
	"fmt"
	"sort"
)

/*
Suppose you have a random list of people standing in a queue. Each person is described by a pair of integers (h, k),where h is the height of the person and k is the number of people in front of this person who have a height greater than or equal to h. Write an algorithm to reconstruct the queue.

Note:
The number of people is less than 1,100.
*/
func reconstructQueue(people [][]int) [][]int {
	var result [][]int
	return result
}

// By is the type of a "less" function that defines the ordering of its [][]int arguments.
type By func(i, j *[]int) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(queue [][]int) {
	qs := &QueueSorter{
		queue: queue,
		by:    by,
	}
	sort.Stable(qs)
}

// QueueSorter collect queue and compare function
type QueueSorter struct {
	queue [][]int
	by    By
}

func (a *QueueSorter) Len() int           { return len(a.queue) }
func (a *QueueSorter) Swap(i, j int)      { a.queue[i], a.queue[j] = a.queue[j], a.queue[i] }
func (a *QueueSorter) Less(i, j int) bool { return a.by(&a.queue[i], &a.queue[j]) }

func main() {
	sortByH := func(i, j *[]int) bool {
		return (*i)[0] < (*j)[0]
	}

	sortByK := func(i, j *[]int) bool {
		return (*i)[1] < (*j)[1]
	}

	queue := [][]int{[]int{4, 2}, []int{1, 6}, []int{9, 77}, []int{3, 10}}

	By(sortByH).Sort(queue)
	fmt.Println(queue)

	By(sortByK).Sort(queue)
	fmt.Println(queue)
}
