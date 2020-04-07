package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Description string
	Price float64
}

type meal struct {
	Meal string
	Items []item
}


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main () {
	menu := []meal{
		meal{
			Meal: "Breakfast",
			Items: []item{
				item{
					Name: "Pancake",
					Description: "Japanese or American Pancakes",
					Price: 1.95,
				},
				item{
					Name: "Fruits",
					Description: "Fresh Fruits",
					Price: 0.95,
				},
			},
		},
		meal{
			Meal: "Lunch",
			Items: []item{
				item{
					Name: "Lasagna",
					Description: "Italian Lasagna",
					Price: 9.99,
				},
				item{
					Name: "Feijoada",
					Description: "Brazilian Feijoada",
					Price: 20.55,
				},
			},
		},
		meal{
			Meal: "Dinner",
			Items: []item{
				item{
					Name: "Pizza",
					Description: "Italian Pizza",
					Price: 9.99,
				},
				item{
					Name: "Sushi",
					Description: "Japanese Sushi",
					Price: 20.55,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, menu)
	if err != nil {
		log.Fatalln(err)
	}
}