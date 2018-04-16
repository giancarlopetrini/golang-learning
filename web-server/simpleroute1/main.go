package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.tmpl.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Route string
		Path  string
	}{
		"/",
		r.URL.Path,
	}
	tpl.Execute(w, data)
}

func dog(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Route string
		Path  string
	}{
		"/dog",
		r.URL.Path,
	}
	tpl.Execute(w, data)
}

func me(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Route string
		Path  string
		Name  string
	}{
		"/me",
		r.URL.Path,
		"giancarlo",
	}
	tpl.Execute(w, data)
}
