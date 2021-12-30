package minimumnumberofkconsecutivebitflips

func minKBitFlips(nums []int, k int) int {
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
