package maxincreasetokeepcityskyline

func maxIncreaseKeepingSkyline(grid [][]int) int {
	rowCount := len(grid)
	if 0 == rowCount {
		return 0
	}
	colCount := len(grid[0])
	if 0 == colCount {
		return 0
	}
	rowMax := make([]int, rowCount)
	colMax := make([]int, colCount)
	for i, v1 := range grid {
		for j, v := range v1 {
			if v > rowMax[i] {
				rowMax[i] = v
			}
			if v > colMax[j] {
				colMax[j] = v
			}
		}
	}
	result := 0
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			// min row, col
			minHeight := func() int {
				if rowMax[i] < colMax[j] {
					return rowMax[i]
				}
				return colMax[j]
			}()
			r := minHeight - grid[i][j]
			result += r
		}
	}
	return result
}
