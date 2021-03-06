package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/KIT-MAMID/fakeapi/apiobjects"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	staticServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/", staticServer)
	router.PathPrefix("/static/").Handler(staticServer)
	router.PathPrefix("/pages/").Handler(staticServer)
	router.Methods("GET").Path("/api/slaves").Name("SlaveIndex").HandlerFunc(apiobjects.SlaveIndex)
	router.Methods("GET").Path("/api/slaves/{slaveId}").Name("SlaveById").HandlerFunc(apiobjects.SlaveById)
	router.Methods("PUT").Path("/api/slaves").Name("SlavePut").HandlerFunc(apiobjects.SlavePut)
	router.Methods("POST").Path("/api/slaves/{slaveId}").Name("SlaveUpdate").HandlerFunc(apiobjects.SlaveUpdate)
	router.Methods("DELETE").Path("/api/slaves/{slaveId}").Name("SlaveDelete").HandlerFunc(apiobjects.SlaveDelete)
	router.Methods("GET").Path("/api/replicasets").Name("ReplicaSetIndex").HandlerFunc(apiobjects.ReplicaSetIndex)
	router.Methods("GET").Path("/api/riskgroups").Name("RiskGroupIndex").HandlerFunc(apiobjects.RiskGroupIndex)

	log.Fatal(http.ListenAndServe(":8080", router))
}
