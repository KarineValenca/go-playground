package main

import (
	"net/http"
	"html/template"
	"log"
)

func index (w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/index.gohtml"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func main () {
	fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", index)
	http.Handle("/pics/", fs)
	http.ListenAndServe(":8080", nil)
}