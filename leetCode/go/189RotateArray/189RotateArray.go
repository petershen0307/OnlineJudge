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
	temp := append(nums[length-rotateNums:], nums[0:length-rotateNums]...)
	for i, v := range temp {
		nums[i] = v
	}
	// for i := 0; i != rotateNums; i++ {
	// 	for j := length - 1; j > 0; j-- {
	// 		nums[j-1], nums[j] = nums[j], nums[j-1]
	// 	}
	// }
}
