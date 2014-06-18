package main

import (
	"github.com/codegangsta/negroni"
	// "github.com/ninnemana/marvelista/api"
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/layout.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Fprint(rw, api.List())
}
