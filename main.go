package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello world")

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
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf(
			"<li class='list-group-item bg-primary text-white'>%s - %s</li>",
			title,
			director,
		)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
