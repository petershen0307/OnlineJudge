package sum3

import (
	"fmt"
	"sort"
)

func addOutput(result [][]int, resultMap map[string]int, a, b, c int) [][]int {
	key := fmt.Sprintf("%d,%d,%d", a, b, c)
	if _, ok := resultMap[key]; !ok {
		result = append(result, []int{a, b, c})
		resultMap[key] = 1
	}
	return result
}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	resultMap := make(map[string]int)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i, a := range nums {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			b := nums[start]
			c := nums[end]
			r := a + b + c
			if 0 == r {
				result = addOutput(result, resultMap, a, b, c)
				start++
				end--
			} else if 0 > r {
				start++
			} else if 0 < r {
				end--
			}
		}
	}
	return result
}
