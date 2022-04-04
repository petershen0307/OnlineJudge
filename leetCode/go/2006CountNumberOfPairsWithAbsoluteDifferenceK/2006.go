package countnumberofpairswithabsolutedifferencek

func countKDifference(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			diff := nums[i] - nums[j]
			if diff == k || diff*-1 == k {
				count++
			}
		}
	}
	return count
}
