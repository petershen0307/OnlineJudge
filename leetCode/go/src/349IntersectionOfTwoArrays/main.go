package main

import (
	"fmt"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	rMap := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				rMap[nums1[i]] = true
				break
			}
		}
	}
	var rSlice []int
	for i := range rMap {
		rSlice = append(rSlice, i)
	}
	return rSlice
}

func main() {
	fmt.Println("result: ", intersection([]int{1, 2, 3, 4}, []int{5, 6, 3, 1}))
}
