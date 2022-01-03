package pairsofsongswithtotaldurationsdivisibleby60

func numPairsDivisibleBy60(time []int) int {
	result := 0
	remainder := [60]int{}
	for _, v := range time {
		remainder[v%60]++
	}
	for _, v := range time {
		r := (60 - (v % 60)) % 60
		// 處理餘數是0, 30
		if r == v%60 {
			remainder[v%60]--
		}
		if remainder[r] > 0 {
			result += remainder[r]
		}
		if r != v%60 {
			remainder[v%60]--
		}
	}
	return result
}
