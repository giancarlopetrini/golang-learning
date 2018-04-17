package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/satori/go.uuid"
)

// User contains user info
type User struct {
	Username string
	Password []byte
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
	http.HandleFunc("/logout", logout)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	if !loggedIn(r) {
		fmt.Println("No user logged in.")
		tpl.ExecuteTemplate(w, "index.tmpl.html", nil)
		return
	}
	c, _ := r.Cookie("session")
	data := getUser(c)
	fmt.Println("User map:	", data)
	tpl.ExecuteTemplate(w, "index.tmpl.html", data)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		un := r.FormValue("username")

		// if username exists already
		if _, ok := userDB[un]; ok {
			data := struct {
				Message string
			}{"Sorry, Username already exists."}
			tpl.ExecuteTemplate(w, "index.tmpl.html", data)
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")

		sID, err := uuid.NewV4()
		if err != nil {
			http.Error(w, "Couldn't create session ID", http.StatusInternalServerError)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		// write to session table
		sessionDB[sID] = un

		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		})

		userDB[un] = User{un, hash, fname, lname}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if r.Method == http.MethodGet {
		if loggedIn(r) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		tpl.ExecuteTemplate(w, "signup.tmpl.html", nil)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	if err := logoutUser(w, r); err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
