package main

import (
	"fmt"
	"sudoku-generator/sudoku"
)

func main() {
	// grid := sudoku.SudokuGrid{}
	// // fill in grid
	// for num := 1; num <= 9; num++ {
	// 	// // pick random spot and place
	// 	row, col := rand.IntN(9), rand.IntN(9)
	// 	for !grid.CellIsValid(row, col, num) {
	// 		row, col = rand.IntN(9), rand.IntN(9)
	// 	}
	// 	grid[row][col] = num
	// }
	grid := sudoku.SudokuGrid{
		{4, 3, 5, 2, 6, 9, 7, 8, 1},
		{6, 8, 2, 5, 7, 1, 4, 9, 3},
		{1, 9, 7, 8, 3, 4, 5, 6, 2},
		{8, 2, 6, 1, 9, 5, 3, 4, 7},
		{3, 7, 4, 6, 8, 2, 9, 1, 5},
		{9, 5, 1, 7, 4, 3, 6, 2, 8},
		{5, 1, 9, 3, 2, 6, 8, 7, 4},
		{2, 4, 8, 9, 5, 7, 1, 3, 6},
		{7, 6, 3, 4, 1, 8, 2, 5, 0},
	}

	fmt.Println(*grid.FillRandom(0))
	// fmt.Println(grid)
}
