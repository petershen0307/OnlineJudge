package minimumabsolutedifference

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	min := 1000001
	for i := 0; i < len(arr)-1; i++ {
		sub := arr[i+1] - arr[i]
		if min > sub {
			min = sub
		}
	}
	result := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		sub := arr[i+1] - arr[i]
		if min == sub {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}
	return result
}
