package jumpgameiii

func canReach(arr []int, start int) bool {
	if start >= len(arr) || start < 0 {
		return false
	}
	if arr[start] < 0 {
		return false
	}
	temp := arr[start]
	arr[start] = -1
	r := canReach(arr, start+temp)
	r = r || canReach(arr, start-temp)
	return r || temp == 0
}
