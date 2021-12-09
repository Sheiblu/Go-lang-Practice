package main

import (
	"fmt"
	"net/http"
)

func index_handeler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello world </h1>")
}
func about_handeler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> This is about page </h1>")
}
func main() {
	fmt.Println("Sheiblu")
	http.HandleFunc("/", index_handeler)
	http.HandleFunc("/about", about_handeler)
	http.ListenAndServe(":8000", nil)
}
