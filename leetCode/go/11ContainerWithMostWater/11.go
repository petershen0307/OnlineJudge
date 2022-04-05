package containerwithmostwater

func maxArea(height []int) int {
	maxValue := 0
	for begin, end := 0, len(height)-1; begin < end; {
		smallVal := height[end]
		xDiff := end - begin
		if height[begin] < height[end] {
			smallVal = height[begin]
			begin++
		} else {
			end--
		}
		if smallVal*xDiff > maxValue {
			maxValue = smallVal * xDiff
		}
	}
	return maxValue
}
