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

func setZeroes2(matrix [][]int) {
	rowC := len(matrix)
	if 0 == rowC {
		return
	}
	colC := len(matrix[0])
	if 0 == colC {
		return
	}
	isFirstRowZero := false
	isFirstColZero := false
	for i := 0; i < rowC; i++ {
		// check first row is zero
		if 0 == matrix[i][0] {
			isFirstRowZero = true
		}
	}
	for j := 0; j < colC; j++ {
		// check first col is zero
		if 0 == matrix[0][j] {
			isFirstColZero = true
		}
	}
	for i := 0; i < rowC; i++ {
		for j := 0; j < colC; j++ {
			// handle row col except first
			if 0 == matrix[i][j] && i != 0 && j != 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < rowC; i++ {
		for j := 1; j < colC; j++ {
			// handle row col except first
			if 0 == matrix[i][0] || 0 == matrix[0][j] {
				matrix[i][j] = 0
			}
		}
	}
	// handle first row and col
	for i := 0; i < rowC; i++ {
		if isFirstRowZero {
			matrix[i][0] = 0
		}
	}
	for j := 0; j < colC; j++ {
		if isFirstColZero {
			matrix[0][j] = 0
		}
	}
}
