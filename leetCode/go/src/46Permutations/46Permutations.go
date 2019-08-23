package permutations

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	result := make([][]int, 0)
	for i, n := range nums {
		// need the copy here
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		tmp = append(tmp[:i], tmp[i+1:]...)
		rr := permute(tmp)
		for _, r := range rr {
			result = append(result, append(r, n))
		}
	}
	return result
}