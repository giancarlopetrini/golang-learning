package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.tmpl.html"))
}

func main() {

	http.HandleFunc("/", startOver)
	http.HandleFunc("/guess", makeGuess)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func startOver(w http.ResponseWriter, r *http.Request) {
	if err := tpl.Execute(w, nil); err != nil {
		log.Fatalln("Could't parse template...")
	}
}

func makeGuess(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//guess := r.Form.Get("letter")
}
