package main

import (
	"fmt"
	"sudoku-generator/sudoku"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		grid := sudoku.SudokuGrid{}
		res := grid.FillRandom(0)
		subtracted := *res.SubtractRandom()
		fmt.Println(subtracted)
		fmt.Println(subtracted.Solve(0))
	}
}
