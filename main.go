package main

import (
	"bytes"
	"fmt"
)

type Grid [9][9]int

func main() {
	grid := Grid{
		// {4, 3, 5, 2, 6, 9, 7, 8, 1},
		// {6, 8, 2, 5, 7, 1, 4, 9, 3},
		// {1, 9, 7, 8, 3, 4, 5, 6, 2},
		// {8, 2, 6, 1, 9, 5, 3, 4, 7},
		// {3, 7, 4, 6, 8, 2, 9, 1, 5},
		// {9, 5, 1, 7, 4, 3, 6, 2, 8},
		// {5, 1, 9, 3, 2, 6, 8, 7, 4},
		// {2, 4, 8, 9, 5, 7, 1, 3, 6},
		// {7, 6, 3, 4, 1, 8, 2, 5, 9},

		{0, 0, 0, 2, 6, 0, 7, 0, 1},
		{6, 8, 0, 0, 7, 0, 0, 9, 0},
		{1, 9, 0, 0, 0, 4, 5, 0, 0},
		{8, 2, 0, 1, 0, 0, 0, 4, 0},
		{0, 0, 4, 6, 0, 2, 9, 0, 0},
		{0, 5, 0, 0, 0, 3, 0, 2, 8},
		{0, 0, 9, 3, 0, 0, 0, 7, 4},
		{0, 4, 0, 0, 5, 0, 0, 3, 6},
		{7, 0, 3, 0, 1, 8, 0, 0, 0},
	}
	sols := solve(grid, 0)
	for _, sol := range sols {
		fmt.Println(sol)
		fmt.Println(sol.isValid())
	}
}

func isValid(grid Grid, r int, c int, num int) bool {
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

func solve(grid Grid, cell int) []Grid {
	solutions := []Grid{}
	if cell == 81 {
		solutions = append(solutions, grid.copy())
		return solutions
	}
	r := cell / 9
	c := cell % 9

	if grid[r][c] != 0 {
		return solve(grid, cell+1)
	}

	for num := 1; num <= 9; num++ {
		if !isValid(grid, r, c, num) {
			continue
		}
		grid[r][c] = num
		solutions = append(solutions, solve(grid, cell+1)...)
		grid[r][c] = 0
	}

	return solutions
}

func (g Grid) copy() Grid {
	copy := Grid{}
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			copy[r][c] = g[r][c]
		}
	}
	return copy
}

func (grid Grid) isValid() bool {
	temp := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			temp, grid[r][c] = grid[r][c], 0
			if !isValid(grid, r, c, temp) {
				grid[r][c] = temp
				return false
			}
			grid[r][c] = temp
		}
	}
	return true
}

func (grid Grid) String() string {
	res := bytes.Buffer{}
	for _, row := range grid {
		for _, num := range row {
			res.WriteString(fmt.Sprintf("%v ", num))
		}
		res.WriteString("\n")
	}
	return res.String()
}
