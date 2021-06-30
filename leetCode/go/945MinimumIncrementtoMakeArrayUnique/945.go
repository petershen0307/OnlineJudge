package minimumincrementtomakearrayunique

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	result := 0
	scan(0, nums, &result)
	return result
}

func scan(start int, nums []int, result *int) {
	if start >= len(nums) {
		return
	}
	nextEntry := start
	for i := start + 1; i < len(nums); i++ {
		if nums[start] == nums[i] {
			(*result)++
			nums[i]++
			if nextEntry == start {
				nextEntry = i
			}
		} else {
			break
		}
	}
	if nextEntry == start {
		nextEntry++
	}
	scan(nextEntry, nums, result)
}
