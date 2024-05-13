package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	templates := template.Must(template.ParseGlob("views/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index", struct{}{})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
