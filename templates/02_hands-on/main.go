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

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []region {
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name: "Babylon",
					Address: "California Street",
					City: "California city",
					Zip: "5555-555",
					Region: "southern",
				},
				hotel{
					Name: "Narnia",
					Address: "4 Street",
					City: "California city",
					Zip: "78587-555",
					Region: "southern",
				},
			},
		},
		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name: "Babylon C",
					Address: "California Street C",
					City: "California city C",
					Zip: "5555-555",
					Region: "central",
				},
				hotel{
					Name: "Narnia",
					Address: "4 Street",
					City: "California city",
					Zip: "78587-555",
					Region: "central",
				},
			},
		},
	}
	// hotels := []hotel{
	// 	hotel{
	// 		Name: "Babylon",
	// 		Address: "California Street",
	// 		City: "California city",
	// 		Zip: "5555-555",
	// 		Region: "Central",
	// 	},
	// 	hotel{
	// 		Name: "Narnia",
	// 		Address: "4 Street",
	// 		City: "California city",
	// 		Zip: "78587-555",
	// 		Region: "southern",
	// 	},
	// }

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}