package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/KIT-MAMID/fakeapi/apiobjects"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/", http.FileServer(http.Dir("./static/")))
	router.Methods("GET").Path("/api/slaves").Name("SlaveIndex").HandlerFunc(apiobjects.SlaveIndex)
	router.Methods("GET").Path("/api/slave/{slaveId}").Name("SlaveById").HandlerFunc(apiobjects.SlaveById)
	router.Methods("GET").Path("/api/replicasets").Name("ReplicaSetIndex").HandlerFunc(apiobjects.ReplicaSetIndex)
	router.Methods("GET").Path("/api/riskgroups").Name("RiskGroupIndex").HandlerFunc(apiobjects.RiskGroupIndex)

	log.Fatal(http.ListenAndServe(":8080", router))
}
