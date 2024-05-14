package main

import (
	"fmt"
	"sudoku-generator/sudoku"
)

func main() {

	// fmt.Println(seed)
	seed := []byte("YrrWpHhd")
	g := sudoku.NewGrid(&seed, true)
	fmt.Println(g)
}
