package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

type person struct {
	Fname string
	Lname string
	Items []string
}

func index(w http.ResponseWriter, r *http.Request) {
	s := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>json simple1</title>
		</head>
		<body>
		<h1>json simple1</h1>
		</body>
		</html>`
	w.Write([]byte(s))
}

func marshal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "Bob",
		Lname: "Smith",
		Items: []string{"item1", "item2", "item3"},
	}

	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func encode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "bob",
		Lname: "smith",
		Items: []string{"item1", "item2", "item3"},
	}

	if err := json.NewEncoder(w).Encode(p1); err != nil {
		log.Println(err)
	}
}
