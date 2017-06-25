package sum3

import (
	"reflect"
	"sort"
)

func checkDuplicate(a [][]int, b []int) bool {
	for _, v := range a {
		if reflect.DeepEqual(v, b) {
			return true
		}
	}
	return false
}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
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
					if !checkDuplicate(result, temp) {
						result = append(result, temp)
					}
				}
			} else if (r == i || r == j) && i != j {
				// r is equal i or j
				if scale, ok := numMap[r]; ok && 2 <= scale {
					temp := []int{i, j, r}
					sort.Ints(temp)
					if !checkDuplicate(result, temp) {
						result = append(result, temp)
					}
				}
			} else {
				if scale, ok := numMap[r]; ok && 3 <= scale {
					temp := []int{i, j, r}
					sort.Ints(temp)
					if !checkDuplicate(result, temp) {
						result = append(result, temp)
					}
				}
			}
		}
	}
	return result
}
