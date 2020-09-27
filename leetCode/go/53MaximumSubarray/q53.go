package q53

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
	sumFunc := func(numbers []int) int {
		sum := 0
		for _, v := range numbers {
			sum += v
		}
		return sum
	}
	isFirst := false
	sum := 0
	for i := range nums {
		for j := i + 1; j <= len(nums); j++ {
			s := sumFunc(nums[i:j])
			if s > sum || !isFirst {
				sum = s
				isFirst = true
			}
		}
	}
	return sum
}
