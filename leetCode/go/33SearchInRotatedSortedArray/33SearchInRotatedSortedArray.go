package searchinrotatedsortedarray

func search(nums []int, target int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			// nums[left] <= nums[mid] 表示在同一個排序過的陣列裡，但是 target 不在 left~mid 之間
			if nums[left] > target && nums[left] <= nums[mid] {
				// 2 3 4 5 6 7 0 1
				// target: 0
				// left: nums[0]=2
				// mid: nums[4]=5
				left = mid + 1
			} else {
				// 2 3 4 5 6 7 0 1
				// target: 4
				// left: nums[0]=2
				// mid: nums[4]=5
				// right: nums[7]=1
				right = mid - 1
			}
		}
		if nums[mid] < target {
			// nums[right] > nums[mid] 表示在同一個排序過的陣列裡，但是 target 不在 mid~right 之間
			if nums[right] < target && nums[right] >= nums[mid] {
				// 6 7 0 1 2 3 4 5
				// target: 7
				// right: nums[7]=5
				// mid: nums[4]=1
				right = mid - 1
			} else {
				// 6 7 0 1 2 3 4 5
				// target: 3
				// right: nums[7]=5
				// mid: nums[4]=1
				// left: nums[0]=6
				left = mid + 1
			}
		}
	}
	return -1
}
