package minimumnumberofkconsecutivebitflips

func minKBitFlips(nums []int, k int) int {
	isFlipped := make([]int, len(nums))
	flipped, result := 0, 0
	for i, v := range nums {
		// 因為 flipped 記錄了前面k個反轉的資訊，所以第k+1個的時候要把第一個是否反轉的資訊扣除
		// N xor 1 xor 1 == N，利用xor兩次會變成原值得特性把第一個反轉的資訊扣除
		if i >= k {
			flipped ^= isFlipped[i-k]
		}
		if flipped == v {
			// i+k超過len(nums), 表示已經沒法再做翻轉了
			if i+k > len(nums) {
				return -1
			}
			isFlipped[i] = 1
			flipped ^= 1
			result++
		}
	}
	return result
}

func minKBitFlips2(nums []int, k int) int {
	one := 0
	zero := 0
	result := 0
	slot := []int{}
	for _, v := range nums {
		if v == 1 {
			one++
		} else {
			zero++
		}
		if len(slot) != k {
			slot = append(slot, v)
		}
		if len(slot) == k {
			// do flip
			if slot[0] == 0 {
				for i := range slot {
					slot[i] ^= 1
				}
				result++
				one, zero = zero, one
			}
			// pop first 1
			one--
			slot = slot[1:]
		}
	}
	if zero != 0 {
		result = -1
	}
	return result
}
