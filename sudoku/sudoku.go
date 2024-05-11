package sudoku

import (
	"bytes"
	"fmt"
	"math/rand/v2"
)

type SudokuGrid [9][9]int

func NewGrid(random bool) SudokuGrid {
	grid := SudokuGrid{}
	if !random {
		return grid
	}
	res := grid.fillRandom(0)
	return *res.subtractRandom()
}

func (grid SudokuGrid) Solve(cell int) []SudokuGrid {
	solutions := []SudokuGrid{}
	if cell == 81 {
		solutions = append(solutions, grid)
		return solutions
	}
	r := cell / 9
	c := cell % 9

	if grid[r][c] != 0 {
		return grid.Solve(cell + 1)
	}

	for num := 1; num <= 9; num++ {
		if !grid.CellIsValid(r, c, num) {
			continue
		}
		grid[r][c] = num
		solutions = append(solutions, grid.Solve(cell+1)...)
		grid[r][c] = 0
	}

	return solutions
}

func (grid SudokuGrid) CellIsValid(r int, c int, num int) bool {
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

// func (g SudokuGrid) Copy() SudokuGrid {
// 	result := SudokuGrid{}
// 	for r := 0; r < len(g); r++ {
// 		copy(result[r][:], g[r][:])
// 	}
// 	return result
// }

func (grid SudokuGrid) GridIsValid() bool {
	temp := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			temp, grid[r][c] = grid[r][c], 0
			if !grid.CellIsValid(r, c, temp) {
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
	for r, row := range grid {
		for c, num := range row {
			if num == 0 {
				res.WriteString("_ ")
			} else {
				res.WriteString(fmt.Sprintf("%v ", num))
			}
			if c == 2 || c == 5 {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n")

		if r == 2 || r == 5 {
			res.WriteString("\n")
		}
	}
	return res.String()
}

func (grid SudokuGrid) fillRandom(cell int) *SudokuGrid {
	if cell == 81 {
		return &grid
	}
	r := cell / 9
	c := cell % 9

	if grid[r][c] != 0 {
		return grid.fillRandom(cell + 1)
	}

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for len(nums) > 0 {
		i := rand.IntN(len(nums))
		num := nums[i]
		nums = append(nums[:i], nums[i+1:]...)
		if !grid.CellIsValid(r, c, num) {
			continue
		}
		grid[r][c] = num
		result := grid.fillRandom(cell + 1)
		if result != nil {
			return result
		}
		grid[r][c] = 0
	}
	return nil
}

func (grid SudokuGrid) subtractRandom() *SudokuGrid {
	// assume filled grid

	remainingCells := 81
	subtractableCells := make([]int, 81)
	for i := range subtractableCells {
		subtractableCells[i] = i
	}

	target := 30 - rand.IntN(5)
	for remainingCells > target && len(subtractableCells) > 0 {
		n := rand.IntN(len(subtractableCells))
		r, c := subtractableCells[n]/9, subtractableCells[n]%9
		temp := grid[r][c]
		grid[r][c] = 0
		if len(grid.Solve(0)) > 1 {
			grid[r][c] = temp
		} else {
			remainingCells--
		}
		subtractableCells = append(subtractableCells[:n], subtractableCells[n+1:]...)
	}

	return &grid
}
