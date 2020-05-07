package main

import (
	"net/http"
	"html/template"
	"github.com/satori/go.uuid"
	"strings"
	"fmt"
	"crypto/sha1"
	"io"
	"os"
	"filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie := getCookie(w, req)
	
	// proccess form submission
	if req.Method == http.MethodPost{
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		//create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fileName := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		//create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fileName)
		//TODO CONTINUOUS HERE
	}

	xs := strings.Split(cookie.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie{
	// look for a cookie with session name
	cookie, err := req.Cookie("cookie")
	if err != nil {
		//if cookie not found, create a cookie
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "cookie",
			Value: id.String(),
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}

func saveImages(w http.ResponseWriter, c *http.Cookie) *http.Cookie{
	p1 := "disnayland.jpg"
	p2 := "atbeach.jpg"
	p3 := "hollywood.jpg"

	s := c.Value

	if !strings.Contains(s, p1) {
		s += "|" + p1
	}

	if !strings.Contains(s, p2) {
		s += "|" + p2
	}

	if !strings.Contains(s, p3) {
		s += "|" + p3
	}

	c.Value = s
	http.SetCookie(w, c)
	
	return c
}