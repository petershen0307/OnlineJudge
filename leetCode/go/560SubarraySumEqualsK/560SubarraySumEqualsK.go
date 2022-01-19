package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	r := 0
	//sum(i,j)=sum(0,j)-sum(0,i), where sum(i,j) represents the sum of all the elements from index i to j-1. Can we use this property to optimize it.
	// V = accumulateSum - k ==> V如果出現在subarraySumCount的key, 表示至少有一組subarray的加總是V
	subarraySumCount := map[int]int{0: 1} // key: sum of subarray, value: subarray count
	for accumulateSum, i := 0, 0; i != len(nums); i++ {
		accumulateSum += nums[i]
		r += subarraySumCount[accumulateSum-k]
		subarraySumCount[accumulateSum] += 1
	}
	return r
}

func subarraySum2(nums []int, k int) int {
	r := 0
	for i := 0; i != len(nums); i++ {
		sum := nums[i]
		if sum == k {
			r++
		}
		for j := i + 1; j != len(nums); j++ {
			sum += nums[j]
			if sum == k {
				r++
			}
		}
	}
	return r
}
