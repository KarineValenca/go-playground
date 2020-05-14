package main

import (
	"./controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"gopkg.in/mgo.v2"
)

func main()  {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	
	r.GET("/user/:id", uc.GetUser)
	// added route
	r.POST("/user", uc.CreateUser)
	// added rout plus parameter
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}