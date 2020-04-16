package main

import (
	"net/http"
	//"io"
	"log"
	"html/template"
)

func index (w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	
	tpl.ExecuteTemplate(w, "index.gohtml", "Say hello to index")
}

func dog (w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	
	tpl.ExecuteTemplate(w, "index.gohtml", "Say hello to dog")
}

func me (w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	
	tpl.ExecuteTemplate(w, "index.gohtml", "Say hello to Karine")
}

var tpl *template.Template
func init () {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main () {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)

}