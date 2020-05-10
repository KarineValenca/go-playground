package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type statusCode struct {
	Code string
	Descrip int
}

type statusTypes []statusCode

func main() {
	var data statusTypes
	rcvd := `[
		{"Code":"StatusOK","Descrip":200},
		{"Code":"StatusMovedPermanently","Descrip":301},
		{"Code":"StatusFound","Descrip":303},
		{"Code":"StatusSeeOther","Descrip":307}]`
	
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range data {
		fmt.Println(v.Code, "-", v.Descrip)
	}
}