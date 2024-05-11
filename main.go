package main

import (
	"log"
	"net/http"

	"sudoku-generator/sudoku"

	"github.com/a-h/templ"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/"):]
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

func main() {
	grid := sudoku.NewGrid(true)
	component := sudokuGrid(grid)
	http.Handle("/", templ.Handler(component))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
