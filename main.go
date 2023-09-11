package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "Lord of the Rings", Director: "Peter Jackson"},
				{
					Title:    "Star Wars: Revenge of the Sith",
					Director: "George Lucas",
				},
			},
		}

		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(
			w,
			"film-list-element",
			Film{Title: title, Director: director},
		)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
