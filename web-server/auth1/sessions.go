package main

import (
	"net/http"

	"github.com/satori/go.uuid"
)

func loggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		return false
	}
	sID, _ := uuid.FromString(c.Value)

	if _, ok := sessionDB[sID]; ok {
		return ok
	}
	// if no username in sessions db, return false
	return false
}

func getUser(c *http.Cookie) User {
	// get sid from, user
	sID, _ := uuid.FromString(c.Value)
	un := sessionDB[sID]
	return userDB[un]

}

func logoutUser(w http.ResponseWriter, r *http.Request) error {
	c, _ := r.Cookie("session")
	c.MaxAge = -1
	tpl.ExecuteTemplate(w, "index.tmpl.html", nil)
	return nil
}
