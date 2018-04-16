package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

// User contains user info
type User struct {
	Username string
	Fname    string
	Lname    string
}

var sessionDB = map[uuid.UUID]string{}
var userDB = map[string]User{}
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/user", user)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.tmpl.html", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "signup.tmpl.html", nil)
}

func user(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "user.tmpl.html", nil)
}
