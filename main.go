package main

import (
	"log"
	"net/http"

	"sudoku-generator/view"

	"github.com/a-h/templ"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/"):]
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

func main() {
	page := view.Base()

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(page))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
