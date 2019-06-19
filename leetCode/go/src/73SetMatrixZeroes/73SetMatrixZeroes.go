package setmatrixzeroes

func setZeroes(matrix [][]int) {
	xLen := len(matrix)
	if 0 == xLen {
		return
	}
	yLen := len(matrix[0])
	if 0 == yLen {
		return
	}
	// individually record x, y point
	// ↓
	xArray := make([]bool, xLen)
	// →
	yArray := make([]bool, yLen)
	for x, vm := range matrix {
		for y, vi := range vm {
			if 0 == vi {
				xArray[x] = true
				yArray[y] = true
			}
		}
	}
	for x, vm := range matrix {
		for y := range vm {
			if xArray[x] || yArray[y] {
				matrix[x][y] = 0
			}
		}
	}
}
