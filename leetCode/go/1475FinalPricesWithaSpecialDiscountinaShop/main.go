package finalpriceswithaspecialdiscountinashop

func finalPrices(prices []int) []int {
	result := make([]int, len(prices))
	for i, v := range prices {
		result[i] = v
		for _, w := range prices[i+1:] {
			if result[i] >= w {
				result[i] = result[i] - w
				break
			}
		}
	}
	return result
}
