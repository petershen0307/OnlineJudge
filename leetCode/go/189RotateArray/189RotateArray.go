package rotatearray

func rotate(nums []int, k int) {
	length := len(nums)
	if length == 0 {
		return
	}
	rotateNums := k % length
	if rotateNums == 0 {
		return
	}
	reverse(nums, 0, length-1)
	reverse(nums, 0, rotateNums-1)
	reverse(nums, rotateNums, length-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
