package minimumcosttoconnectsticks

import "sort"

// 1 <= sticks.length <= 10^4
// 1 <= sticks[i] <= 10^4
func connectSticks(sticks []int) int {
	cost := 0
	sort.Ints(sticks)
	for len(sticks) > 1 {
		newStick := sticks[0] + sticks[1]
		// pop first 2 sticks
		sticks = sticks[2:]
		cost += newStick
		for i, s := range sticks {
			if newStick < s {
				sticks = append(sticks, 0)
				copy(sticks[i+1:], sticks[i:])
				sticks[i] = newStick
				break
			}
			if i == len(sticks)-1 {
				sticks = append(sticks, newStick)
				break
			}
		}
	}
	return cost
}
