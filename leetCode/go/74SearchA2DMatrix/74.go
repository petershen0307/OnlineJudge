package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	found := false
	for _, list := range matrix {
		for i := 0; i < len(list) && target >= list[0] && target <= list[len(list)-1]; i++ {
			if list[i] == target {
				found = true
				break
			}
		}
	}
	return found
}
