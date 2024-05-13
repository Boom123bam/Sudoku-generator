package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sudoku-generator/sudoku"
)

type FormData struct {
	Pages          int
	SudokusPerPage int
}

type FormConstraints struct {
	MaxPages              int
	SudokusPerPageOptions []int
}

var formConstraints = FormConstraints{
	MaxPages:              10,
	SudokusPerPageOptions: []int{1, 2, 6},
}

type Page struct {
	Sudokus []sudoku.SudokuGrid
}

type SudokusData struct {
	Pages []Page
}

func (formData FormData) IsValid() bool {
	if formData.Pages < 0 || formData.Pages > formConstraints.MaxPages {
		return false
	}
	for _, n := range formConstraints.SudokusPerPageOptions {
		if n == formData.SudokusPerPage {
			return true
		}
	}
	return false
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	templates := template.New("")
	templates.Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }})

	_, err := templates.ParseGlob("views/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index", formConstraints)
	})

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		formData := FormData{}
		err := error(nil)
		formData.Pages, err = strconv.Atoi(r.Form.Get("pages"))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		formData.SudokusPerPage, err = strconv.Atoi(r.Form.Get("sudokusPerPage"))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if !formData.IsValid() {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		fmt.Println(formData)

		sudokusData := SudokusData{}

		for page := 0; page < formData.Pages; page++ {
			page := Page{}
			for i := 0; i < formData.SudokusPerPage; i++ {
				sudoku := sudoku.NewGrid(true)
				page.Sudokus = append(page.Sudokus, sudoku)
			}
			sudokusData.Pages = append(sudokusData.Pages, page)
		}

		templates.ExecuteTemplate(w, "sudokus", sudokusData)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
