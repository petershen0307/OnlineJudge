package greatestsumdivisiblebythree

import "slices"

func maxSumDivThree(nums []int) int {
	/* 先把所有數字相加 sum，並且根據除以3的餘數分成3組
	   如果sum除以3餘數為0 解答
	   如果sum除以3餘數為1 嘗試移除餘數為1的最小數字，或是移除餘數為2的最小兩個數字
	   如果sum除以3餘數為2 移除餘數為2的最小數字，或移除餘數為1的兩個最小數字
	*/
	slices.Sort(nums)
	sum := 0
	mod_1 := []int{}
	mod_2 := []int{}
	for _, n := range nums {
		sum += n
		switch n % 3 {
		case 1:
			mod_1 = append(mod_1, n)
		case 2:
			mod_2 = append(mod_2, n)
		}
	}
	switch sum % 3 {
	case 1:
		sum_1 := 0
		sum_2 := 0
		if len(mod_1) >= 1 {
			sum_1 = sum - mod_1[0]
		}
		if len(mod_2) >= 2 {
			sum_2 = sum - mod_2[0] - mod_2[1]
		}
		sum = max(sum_1, sum_2)
	case 2:
		sum_1 := 0
		sum_2 := 0
		if len(mod_1) >= 2 {
			sum_1 = sum - mod_1[0] - mod_1[1]
		}
		if len(mod_2) >= 1 {
			sum_2 = sum - mod_2[0]
		}
		sum = max(sum_1, sum_2)
	}
	return sum
}
