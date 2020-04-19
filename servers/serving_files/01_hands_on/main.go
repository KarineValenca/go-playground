package main

import (
	"net/http"
	"io"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func foo (w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog (w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func chien (w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.png")
}

func main () {
	
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.png", chien)
	http.ListenAndServe(":8080", nil)
}