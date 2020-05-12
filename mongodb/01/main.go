package main

import (
	"./controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main()  {
	r := httprouter.New()
	uc := controllers.NewUserController()
	
	r.GET("/user/:id", uc.GetUser)
	// added route
	r.POST("/user", uc.CreateUser)
	// added rout plus parameter
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}