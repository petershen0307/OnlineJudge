package containerwithmostwater

func maxArea(height []int) int {
	xCount := len(height) - 1
	maxValue := 0
	for begin, end := 0, len(height)-1; xCount > 0; xCount-- {
		smallVal := height[end]
		if height[begin] < height[end] {
			smallVal = height[begin]
			begin++
		} else {
			end--
		}
		if smallVal*xCount > maxValue {
			maxValue = smallVal * xCount
		}
	}
	return maxValue
}
