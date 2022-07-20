package main

import (
	"html/template"
	"net/http"
)

type ToDo struct {
	Name     string
	College  string
	Keywords map[string]string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	todos := ToDo{Name: "test", College: "testeset", Keywords: map[string]string{"Gitlab": "Gitlab", "Other": "Other"}}

	t, err := template.ParseFiles("static/index.htm",
		"static/html/nav.htm",
		"static/html/footer.htm",
		"static/html/card.htm",
		"static/html/simplecard.htm",
		"static/html/new.htm",
	)
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, todos)
	if err != nil {
		panic(err)
	}
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	todos := ToDo{Name: "test", College: "testeset"}

	t, _ := template.ParseFiles("static/home.htm",
		"static/html/nav.htm",
		"static/html/footer.htm",
		"static/html/card.htm")
	err := t.Execute(w, todos)
	if err != nil {
		panic(err)
	}
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/login.htm",
		"static/html/nav.htm",
		"static/html/footer.htm",
		"static/html/card.htm")
	err := t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
func testHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/test.htm",
		"static/html/nav.htm",
		"static/html/footer.htm")
	err := t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/test", testHandler)

	http.ListenAndServe(":3000", nil)
}
