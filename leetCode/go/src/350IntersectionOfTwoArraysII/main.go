package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	nums2Tmp := make([]int, len(nums2))
	copy(nums2Tmp, nums2)
	countMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2Tmp); {
			if nums1[i] == nums2Tmp[j] {
				countMap[nums1[i]]++
				nums2Tmp = append(nums2Tmp[:j], nums2Tmp[j+1:]...)
				break
			} else {
				j++
			}
		}
	}
	var rSlice []int
	for key, count := range countMap {
		for i := 0; i < count; i++ {
			rSlice = append(rSlice, key)
		}
	}
	return rSlice
}

func main() {
	fmt.Println("result:", intersect([]int{1, 2, 3, 4, 2, 2}, []int{2, 2, 5, 6, 3, 1}))
	fmt.Println("result:", intersect([]int{}, []int{2, 2, 5, 6, 3, 1}))
	fmt.Println("result:", intersect([]int{2, 3, 2, 3, 2}, []int{2, 2, 5, 6, 3, 3}))
}
