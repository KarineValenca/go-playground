package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type user struct {
	UserName string
	Password []byte
	First string
	Last string
	Role string
}

type session struct {
	un string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time
const sessionLength int = 300

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		// checks if username is already taken
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "UserName already taken", http.StatusForbidden)
			return
		}

		//create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		u := user{un, bs, f, l, r}
		dbUsers[un] = u
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	c, _ := req.Cookie("session")
	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	

	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

