package shortestdistancetotargetcolor

func shortestDistanceColor(colors []int, queries [][]int) []int {
	colorIndexMap := make(map[int][]int, 0)
	colorIndexMap[1] = make([]int, 0)
	colorIndexMap[2] = make([]int, 0)
	colorIndexMap[3] = make([]int, 0)
	// put colors to map
	for i, c := range colors {
		colorIndexMap[c] = append(colorIndexMap[c], i)
	}
	minSearch := func(index []int, color int) int {
		for {
			pivot := len(index) / 2
			if len(index) <= 3 {
				m := 60000
				for _, i := range index {
					r := color - i
					if r < 0 {
						r *= -1
					}
					if m > r {
						m = r
					}
				}
				if m == 60000 {
					m = -1
				}
				return m
			}
			if color > index[pivot] {
				index = index[pivot:]
			} else {
				index = index[:pivot+1]
			}
		}
	}
	result := make([]int, 0)
	for _, q := range queries {
		result = append(result, minSearch(colorIndexMap[q[1]], q[0]))
	}
	return result
}
