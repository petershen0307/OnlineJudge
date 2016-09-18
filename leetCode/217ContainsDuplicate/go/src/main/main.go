package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{}))
	fmt.Println(containsDuplicate([]int{1, 3, 5, 4, 2}))
	fmt.Println(containsDuplicate([]int{1, 3, 5, 4, 2, 55, 66, 3}))
	fmt.Println(containsDuplicate([]int{1, 3, 5, 4, 2, 55, 66, 3, 1}))
}
