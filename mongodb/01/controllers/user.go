package controllers

import (
	"encoding/json"
	"fmt"
	"../models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)



type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name: "James Bond",
		Gender: "male",
		Age: 32,
		Id: p.ByName("id"),
	}

	//marshal into json
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	//Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//composite literak - type and curly brace
	u := models.User{}

	// encode/decode for sending;receiving json to/from a stream
	json.NewDecoder(r.Body).Decode(&u)

	// change id
	u.Id = "007"

	// marshal/unmarshal for having json assigned to a variable
	uj, _ := json.Marshal(u)

	//Write content-type, status code, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//TODO: write code to delete user
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}