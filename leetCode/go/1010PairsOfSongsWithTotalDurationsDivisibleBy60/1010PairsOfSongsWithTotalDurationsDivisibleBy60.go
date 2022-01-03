package pairsofsongswithtotaldurationsdivisibleby60

func numPairsDivisibleBy60_2(time []int) int {
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

func numPairsDivisibleBy60(time []int) int {
	result := 0
	remainder := [60]int{}
	for _, v := range time {
		remainder[v%60]++
	}
	for i := 1; i < 30; i++ {
		result += remainder[i] * remainder[len(remainder)-i]
	}
	// 0
	result += remainder[0] * (remainder[0] - 1) / 2
	// 30
	result += remainder[30] * (remainder[30] - 1) / 2
	return result
}
