package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/ninnemana/marvelista/api"
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
	fmt.Fprint(rw, api.List())
}
