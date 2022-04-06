package sumwithmultiplicity

func threeSumMulti(arr []int, target int) int {
	result := 0
	// 0 <= arr[i] <= 100
	nums := [101]int{}
	for _, a := range arr {
		nums[a]++
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			k := target - i - j
			if k > 100 || k < 0 {
				continue
			}
			// 維持 i <= j <= k 這樣的順序，避免重複計算 k<=i<=j 這種情況
			// 因為 i,j 從0開始遞增，會涵蓋到 k<=i<=j 這種組合
			// 0 <= i <= j <= k <= 100
			if i == j && j == k {
				result += nums[i] * (nums[j] - 1) * (nums[k] - 2) / 6 // 6 = 3!
			} else if i == j {
				result += (nums[i] * (nums[j] - 1)) / 2 * nums[k] // 2 = 2!
			} else if j < k {
				result += nums[i] * nums[j] * nums[k]
			}

		}
	}
	return result % (1e9 + 7)
}
