package main

import (
	"fmt"
	"net/http"
	"strconv"
)


func main() {

	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("nice-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie {
			Name: "nice-cookie",
			Value: "0",
		}
	}
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		fmt.Println(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	
	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "Cookie writen - check your browser")
	fmt.Fprintln(w, "in chrome go to: dev tools/ application / cookies")
}