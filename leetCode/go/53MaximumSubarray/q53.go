package q53

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := nums[0]
	currentSum := nums[0]
	for _, v := range nums[1:] {
		if v > currentSum+v {
			currentSum = v
		} else {
			currentSum += v
		}
		if currentSum > sum {
			sum = currentSum
		}
	}
	return sum
}
