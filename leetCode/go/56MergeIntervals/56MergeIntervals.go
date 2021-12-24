package mergeintervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// sort at initial
	sort.Sort(interval(intervals))
	// record current interval
	i, j := -1, -1
	result := [][]int{}
	for _, pair := range intervals {
		// first initial
		if i == j && i == -1 {
			i, j = pair[0], pair[1]
			continue
		}
		// consider consecutive same value case {1,4}, {4,6}
		if pair[0] <= j {
			// consider case {1,5},{2,4}
			if j < pair[1] {
				j = pair[1]
			}
		} else {
			result = append(result, []int{i, j})
			i, j = pair[0], pair[1]
		}
	}
	if i != -1 && j != -1 {
		result = append(result, []int{i, j})
	}
	return result
}

type interval [][]int

func (x interval) Len() int {
	return len(x)
}

func (x interval) Less(i, j int) bool {
	return x[i][0] < x[j][0]
}

func (x interval) Swap(i, j int) {
	x[i][0], x[j][0] = x[j][0], x[i][0]
	x[i][1], x[j][1] = x[j][1], x[i][1]
}
