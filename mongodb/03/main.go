package main

import (
	"net/http"
	"./controllers"
)

func main() {
	uc := controllers.NewUsersController()
	http.HandleFunc("/", uc.Index)
	http.HandleFunc("/bar", uc.Bar)
	http.HandleFunc("/signup", uc.Signup)
	http.HandleFunc("/login", uc.Login)
	http.HandleFunc("/logout", uc.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}




