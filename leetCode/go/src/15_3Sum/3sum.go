package sum3

import (
	"fmt"
	"sort"
	"strings"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	resultKey := make(map[string]int)
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}
	for indexI, i := range nums {
		for indexJ, j := range nums {
			if indexI >= indexJ {
				continue
			}
			r := -1 * (i + j)
			if r != i && r != j {
				// r is different with i and j
				if _, ok := numMap[r]; ok {
					temp := []int{i, j, r}
					sort.Ints(temp)
					tempKey := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(temp)), ","), "[]")
					if _, ok := resultKey[tempKey]; !ok {
						result = append(result, temp)
						resultKey[tempKey] = 1
					}
				}
			} else if (r == i || r == j) && i != j {
				// r is equal i or j
				if scale, ok := numMap[r]; ok && 2 <= scale {
					temp := []int{i, j, r}
					sort.Ints(temp)
					tempKey := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(temp)), ","), "[]")
					if _, ok := resultKey[tempKey]; !ok {
						result = append(result, temp)
						resultKey[tempKey] = 1
					}
				}
			} else {
				if scale, ok := numMap[r]; ok && 3 <= scale {
					temp := []int{i, j, r}
					sort.Ints(temp)
					tempKey := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(temp)), ","), "[]")
					if _, ok := resultKey[tempKey]; !ok {
						result = append(result, temp)
						resultKey[tempKey] = 1
					}
				}
			}
		}
	}
	return result
}
