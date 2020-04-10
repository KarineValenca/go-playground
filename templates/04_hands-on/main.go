package main

import (
	"log"
	"os"
	"text/template"
)

type stocks struct {
	Date string
	OpenPrice float64
}