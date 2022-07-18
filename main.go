package main

import (
	"html/template"
	"net/http"
)

type ToDo struct {
	Name     string
	College  string
	Keywords []string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	todos := ToDo{Name: "test", College: "testeset", Keywords: []string{"here", "here2"}}

	t, _ := template.ParseFiles("static/index.html",
		"static/html/nav.html",
		"static/html/footer.html",
		"static/html/card.html")
	err := t.Execute(w, todos)
	if err != nil {
		panic(err)
	}
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	todos := ToDo{Name: "test", College: "testeset"}

	t, _ := template.ParseFiles("static/home.html",
		"static/html/nav.html",
		"static/html/footer.html",
		"static/html/card.html")
	err := t.Execute(w, todos)
	if err != nil {
		panic(err)
	}
}
func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/home", homeHandler)

	http.ListenAndServe(":3000", nil)
}
