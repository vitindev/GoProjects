package main

import (
	"html/template"
	"log"
	"net/http"
)

/*
	HTML - HYPER TEXT MARKUP LANGUAGE
*/

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		usuario := usuario{"Victor", "victtorhugo.mendes@gmail.com"}

		templates.ExecuteTemplate(w, "home.html", usuario)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}
