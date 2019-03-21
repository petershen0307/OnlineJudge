package twosum

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for ii, vv := range nums[i+1:] {
			if (v + vv) == target {
				return []int{i, (i + ii + 1)}
			}
		}
	}
	return []int{}
}
