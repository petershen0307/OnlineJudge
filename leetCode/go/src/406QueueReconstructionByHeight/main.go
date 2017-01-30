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
	greatThanByH := func(i, j *[]int) bool {
		return (*i)[0] > (*j)[0]
	}

	lessThanByK := func(i, j *[]int) bool {
		return (*i)[1] < (*j)[1]
	}
	By(lessThanByK).StableSort(people)
	By(greatThanByH).StableSort(people)

	var result [][]int
	for i := 0; i < len(people); i++ {
		// inser to result
		fmt.Println(people, i)
		temp := make([][]int, len(result)-people[i][1])
		copy(temp, result[people[i][1]:len(result)])
		result = append(result[:people[i][1]], people[i])
		result = append(result, temp...)
	}

	return result
}

// By is the type of a "less" function that defines the ordering of its [][]int arguments.
type By func(i, j *[]int) bool

// StableSort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) StableSort(queue [][]int) {
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
	slices := [][]int{[]int{5, 0}, []int{7, 0}, []int{6, 1}, []int{5, 2}, []int{7, 1}, []int{4, 4}}
	fmt.Println(slices)
	slices = reconstructQueue(slices)
	fmt.Println(slices)
}
