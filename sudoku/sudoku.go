package sudoku

import (
	"bytes"
	"fmt"
)

type SudokuGrid [9][9]int

func (grid SudokuGrid) Solve(cell int) []SudokuGrid {
	solutions := []SudokuGrid{}
	if cell == 81 {
		solutions = append(solutions, grid.copy())
		return solutions
	}
	r := cell / 9
	c := cell % 9

	if grid[r][c] != 0 {
		return grid.Solve(cell + 1)
	}

	for num := 1; num <= 9; num++ {
		if !grid.isValidCell(r, c, num) {
			continue
		}
		grid[r][c] = num
		solutions = append(solutions, grid.Solve(cell+1)...)
		grid[r][c] = 0
	}

	return solutions
}

func (grid SudokuGrid) isValidCell(r int, c int, num int) bool {
	// horizontal
	for _, n := range grid[r] {
		if n == num {
			return false
		}
	}
	// vertical
	for i := 0; i < len(grid); i++ {
		if grid[i][c] == num {
			return false
		}
	}
	// 3x3
	startRow := r - r%3
	startCol := c - c%3
	for row := startRow; row < startRow+3; row++ {
		for col := startCol; col < startCol+3; col++ {
			if grid[row][col] == num {
				return false
			}
		}
	}

	return true
}

func (g SudokuGrid) copy() SudokuGrid {
	copy := SudokuGrid{}
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			copy[r][c] = g[r][c]
		}
	}
	return copy
}

func (grid SudokuGrid) isValid() bool {
	temp := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			temp, grid[r][c] = grid[r][c], 0
			if !grid.isValidCell(r, c, temp) {
				grid[r][c] = temp
				return false
			}
			grid[r][c] = temp
		}
	}
	return true
}

func (grid SudokuGrid) String() string {
	res := bytes.Buffer{}
	for _, row := range grid {
		for _, num := range row {
			res.WriteString(fmt.Sprintf("%v ", num))
		}
		res.WriteString("\n")
	}
	return res.String()
}
