package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name string
	Address string
	City string
	Zip string
	Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{
			Name: "Babylon",
			Address: "California Street",
			City: "California city",
			Zip: "5555-555",
			Region: "Central",
		},
		hotel{
			Name: "Narnia",
			Address: "4 Street",
			City: "California city",
			Zip: "78587-555",
			Region: "southern",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}