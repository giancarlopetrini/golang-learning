package main

import (
	"errors"
	"fmt"
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
	fmt.Printf("Current userDB map:\t %v\n", userDB)
	return userDB[un]

}

func logoutUser(r *http.Request) error {
	c, err := r.Cookie("session")
	fmt.Println("Cookie passed to logoutUser:	", c)
	if err == http.ErrNoCookie {
		return errors.New("logged out - No user logged in")
	}
	c.MaxAge = -1
	return nil
}
