package sudoku

import (
	"testing"
)

func TestSolve(t *testing.T) {
	grid := SudokuGrid{
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
	expectedSolution := SudokuGrid{
		{4, 3, 5, 2, 6, 9, 7, 8, 1},
		{6, 8, 2, 5, 7, 1, 4, 9, 3},
		{1, 9, 7, 8, 3, 4, 5, 6, 2},
		{8, 2, 6, 1, 9, 5, 3, 4, 7},
		{3, 7, 4, 6, 8, 2, 9, 1, 5},
		{9, 5, 1, 7, 4, 3, 6, 2, 8},
		{5, 1, 9, 3, 2, 6, 8, 7, 4},
		{2, 4, 8, 9, 5, 7, 1, 3, 6},
		{7, 6, 3, 4, 1, 8, 2, 5, 9},
	}
	sols := grid.Solve(0)

	if len(sols) != 1 {
		t.Errorf("expected 1 solution got %d", len(sols))
	}
	solution := sols[0]
	for r, row := range solution {
		for c := range row {
			if solution[r][c] != expectedSolution[r][c] {
				t.Errorf("solutions dont match at r:%d c:%d\nexpected:\n%s\ngot:\n%s", r, c, expectedSolution, solution)
			}
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	grid := SudokuGrid{
		{0, 0, 0, 5, 7, 3, 9, 0, 6},
		{3, 9, 6, 1, 0, 0, 0, 0, 5},
		{1, 5, 7, 9, 0, 0, 3, 0, 0},
		{0, 0, 0, 6, 9, 2, 7, 0, 0},
		{0, 0, 3, 0, 1, 0, 0, 0, 9},
		{0, 2, 8, 0, 0, 4, 0, 5, 1},
		{0, 7, 2, 0, 0, 0, 0, 0, 0},
		{5, 0, 0, 2, 8, 7, 0, 6, 3},
		{0, 0, 1, 0, 6, 0, 0, 0, 7},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = grid.Solve(0)
	}
}
