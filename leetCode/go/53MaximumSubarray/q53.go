package q53

import "sync"

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
	rChan := make(chan int, len(nums))
	var wg sync.WaitGroup
	wg.Add(len(nums))
	for i := range nums {
		tempNums := nums[i:]
		go func(numbers []int) {
			defer wg.Done()
			isFirst := false
			sum := 0
			for j := range numbers {
				for k := j + 1; k <= len(numbers); k++ {
					s := sumFunc(numbers[j:k])
					if s > sum || !isFirst {
						isFirst = true
						sum = s
					}
				}
			}
			rChan <- sum
		}(tempNums)
	}
	wg.Wait()
	sum := 0
	for i := 0; i < len(nums); i++ {
		s := <-rChan
		if i == 0 || s > sum {
			sum = s
		}
	}
	return sum
}
