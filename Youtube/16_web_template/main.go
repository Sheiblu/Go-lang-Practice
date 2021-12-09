package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var testTemplate *template.Template

type Auther struct {
	Name string
	Age  int
}
type NewsAggPage struct {
	Title   string
	News    string
	Authers []Auther
}

func index_handeler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello world </h1>")
}

func about_handeler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> This is about page </h1>")
}

func newsAgg_handeler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Tutorial Title ",
		News: "Some news"}

	p.Authers = append(p.Authers, Auther{"Sheiblu", 25})
	p.Authers = append(p.Authers, Auther{"Shihab", 22})
	p.Authers = append(p.Authers, Auther{"Ali", 30})

	t, _ := template.ParseFiles("basictemplate.html")
	fmt.Println(t.Execute(w, p))
}

func main() {
	fmt.Println("Sheiblu")
	http.HandleFunc("/", index_handeler)
	http.HandleFunc("/about", about_handeler)
	http.HandleFunc("/template", newsAgg_handeler)
	http.ListenAndServe(":8000", nil)
}
