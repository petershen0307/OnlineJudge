package q1170

func numSmallerByFrequency(queries []string, words []string) []int {
	wMap := make([]int, len(words))
	result := make([]int, len(queries))
	for i, str := range words {
		wMap[i] = f(str)
	}
	for i, str := range queries {
		fq := f(str)
		for _, v := range wMap {
			if fq < v {
				result[i]++
			}
		}
	}
	return result
}

func f(str string) int {
	rMap := [26]int{}
	for _, r := range str {
		rMap[r-'a']++
	}
	for _, v := range rMap {
		if v != 0 {
			return v
		}
	}
	return 0
}
