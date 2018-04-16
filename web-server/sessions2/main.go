package main

import (
	"fmt"
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

var tpl *template.Template
var userDB = make(map[string]User)
var sessionDB = make(map[uuid.UUID]string)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", sessionHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c, err := r.Cookie("session")
		if err == http.ErrNoCookie {
			tpl.ExecuteTemplate(w, "index.tmpl.html", nil)
			return
		} else if err != nil {
			log.Fatalf("Couldn't get index tample: %s\n", err)
			return
		} else {
			// valid session cookie
			sID, err := uuid.FromString(c.Value)
			if err != nil {
				fmt.Println("Couldn't get SID", err)
				return
			}

			// if sessiion ID exists, but no matching user, delete session
			if _, ok := sessionDB[sID]; !ok {
				c.MaxAge = -1
				tpl.ExecuteTemplate(w, "index.tmpl.html", nil)
				return
			}

			un := sessionDB[sID]

			data := struct {
				User
				SID string
			}{
				userDB[un],
				sID.String(),
			}

			tpl.ExecuteTemplate(w, "user.tmpl.html", data)
			fmt.Printf("Session ID:\t %s\n", c.Value)
			fmt.Println("Complete User Map:", userDB)
			fmt.Println("Complete Session ID Map:", sessionDB)
		}
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")

		if _, ok := userDB[un]; ok {
			data := "Username already in use, please choose another."
			tpl.ExecuteTemplate(w, "index.tmpl.html", data)
			return
		}

		fmt.Println("Posted...")
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)

		userDB[un] = User{un, fname, lname}
		sessionDB[sID] = un

		data := struct {
			User
			SID string
		}{
			userDB[un],
			sID.String(),
		}
		tpl.ExecuteTemplate(w, "user.tmpl.html", data)
	}

}
