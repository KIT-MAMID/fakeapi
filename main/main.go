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
	router.HandleFunc("/api/slaves", apiobjects.SlaveIndex)
	router.HandleFunc("/api/slave/{slaveId}", apiobjects.SlaveById)
	router.HandleFunc("/api/replicasets", apiobjects.ReplicaSetIndex)
	router.HandleFunc("/api/riskgroups", apiobjects.RiskGroupIndex)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
}