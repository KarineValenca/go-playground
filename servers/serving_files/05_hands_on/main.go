package main

import (
	"log"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/index.gohtml"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}