package countbinarysubstrings

func countBinarySubstrings(s string) int {
	// 初始值給 1 因為 index 是從1開始, 先把 index 0的數值填入 groups
	// 相同的值就可以直接利用 groups 最後一個值加加
	groups := []int{1}
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			groups = append(groups, 1)
		} else {
			groups[len(groups)-1]++
		}
	}
	answer := 0
	// 至少要有兩個值在 groups 才會滿足有相同 0,1 子字串
	for i := 1; i < len(groups); i++ {
		if groups[i-1] > groups[i] {
			answer += groups[i]
		} else {
			answer += groups[i-1]
		}
	}
	return answer
}
