package main

import (
	"net/http"
	"fmt"
	"html"
	"log"

	"github.com/gorilla/mux"
	"../apiobjects"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/slaves", apiobjects.SlaveIndex)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
}