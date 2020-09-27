package q1186

func maximumSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sum := arr[0]
	currentSum := arr[0]
	skipOneSum := arr[0]
	// start from index 1
	for _, v := range arr[1:] {
		skipLastPlusCurrent := skipOneSum + v
		skipCurrent := currentSum
		noSkip := currentSum + v
		// find the max one
		skipOneSum = max(noSkip, max(skipLastPlusCurrent, skipCurrent))
		currentSum = max(noSkip, v)
		sum = max(skipOneSum, max(sum, currentSum))
	}
	return sum
}
