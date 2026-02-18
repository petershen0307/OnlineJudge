package binarynumberwithalternatingbits

func hasAlternatingBits(n int) bool {
	// 判斷是否連續兩個 bit 皆為同數值
	last := -1
	for ; n > 0; n = n >> 1 {
		if last < 0 {
			last = n & 1
			continue
		}
		if last == n&1 {
			return false
		}
		last = n & 1
	}
	return true
}
