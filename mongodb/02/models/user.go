  
package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id     string `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}

func StoreUser(mu map[string]User) {
	f, err := os.Create("data")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(mu)
}

func GetUsers() map[string]User {
	m := make(map[string]User)
	
	f, err := os.Open("./data")
	if err != nil {
		fmt.Println(err)
		return m
	}

	defer f.Close()
	
	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}