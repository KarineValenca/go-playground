package controllers

import (
	"encoding/json"
	"fmt"
	"../models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//grab id
	id := p.ByName("id")

	// verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//ObjectIdHex return an ObjectId from the provided hex representation
	oid := bson.ObjectIdHex(id)

	// composite literal
	u := models.User{}

	// fetch user
	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
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

	// create bson ID
	u.Id = bson.NewObjectId()

	//store the user in mongodb
	uc.session.DB("go-web-dev-db").C("users").Insert(u)

	// marshal/unmarshal for having json assigned to a variable
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	//Write content-type, status code, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Delete user
	if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}