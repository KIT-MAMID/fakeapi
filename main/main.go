package main

import (
	"net/http"
	"fmt"
	"html"
	"log"

	"github.com/gorilla/mux"
	"github.com/KIT-MAMID/fakeapi/apiobjects"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/", Index)
	router.Methods("GET").Path("/api/slaves").Name("SlaveIndex").HandlerFunc(apiobjects.SlaveIndex)
	router.Methods("GET").Path("/api/slave/{slaveId}").Name("SlaveById").HandlerFunc(apiobjects.SlaveById)
	router.Methods("GET").Path("/api/replicasets").Name("ReplicaSetIndex").HandlerFunc(apiobjects.ReplicaSetIndex)
	router.Methods("GET").Path("/api/riskgroups").Name("RiskGroupIndex").HandlerFunc(apiobjects.RiskGroupIndex)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
}