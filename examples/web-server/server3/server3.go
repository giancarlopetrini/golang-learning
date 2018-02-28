package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler ranges over the structs found in the request and prints them out accordingly
func handler(w http.ResponseWriter, r *http.Request) {
	// we take w (ResponsWriter and format printing of the method, url, and protocol (HTTP/1 here))
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
	}
	fmt.Fprintf(w, "Host = %q]\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q} = %q\n", k, v)
	}
}
