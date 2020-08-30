/*Package IslandPerimeter You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.
Grid cells are connected horizontally/vertically (not diagonally).
The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).
The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1.
The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.*/
package IslandPerimeter

import "log"

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for rowI, rowValue := range grid {
		for colI, value := range rowValue {
			log.Println("row", rowI, ", colI", colI)
			if 0 == value {
				continue
			}
			// up
			if rowI-1 < 0 {
				perimeter++
			} else if 0 == grid[rowI-1][colI] {
				perimeter++
			}
			//down
			if rowI+1 >= len(grid) {
				perimeter++
			} else if 0 == grid[rowI+1][colI] {
				perimeter++
			}
			//left
			if colI-1 < 0 {
				perimeter++
			} else if 0 == grid[rowI][colI-1] {
				perimeter++
			}
			//right
			if colI+1 >= len(rowValue) {
				perimeter++
			} else if 0 == grid[rowI][colI+1] {
				perimeter++
			}
		}
	}
	return perimeter
}
