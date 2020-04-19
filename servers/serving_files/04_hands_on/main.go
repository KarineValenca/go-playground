package main

import (
	"net/http"
	"html/template"
	"log"
)

func dog(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/index.gohtml"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}