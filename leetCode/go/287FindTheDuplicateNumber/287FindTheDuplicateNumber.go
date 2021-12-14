package findtheduplicatenumber

func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		// not support case
		return 0
	}
	for _, v := range nums {
		if v > len(nums) || v <= 0 {
			// not support case
			return 0
		}
	}
	// initial
	n := nums[0]
	nums[0] *= -1
	for nums[n] > 0 {
		t := nums[n]
		nums[n] *= -1
		n = t
	}
	return n
}
