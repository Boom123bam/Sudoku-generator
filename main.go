package main

import (
	"fmt"
	"sudoku-generator/sudoku"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		grid := sudoku.NewGrid(true)
		fmt.Println(grid)
		fmt.Println(grid.Solve(0))
	}
}
