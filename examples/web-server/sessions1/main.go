package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var viewCount int

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	viewCount++
	http.SetCookie(w, &http.Cookie{
		Name:  "increment",
		Value: strconv.Itoa(viewCount),
	})
	fmt.Fprint(w, "Cookie Created and Incremented")
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("increment")
	if err != nil {
		http.NotFoundHandler()
	}
	fmt.Fprintf(w, "Page has been visited %s time(s).", cookie.Value)

}
