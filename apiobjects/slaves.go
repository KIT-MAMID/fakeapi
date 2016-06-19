package apiobjects

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

type Slave struct {
	Id uint			`json:"id"`
	Hostname string		`json:"hostname"`
	Port uint		`json:"slave_port"`
}

func getSlaves() []Slave {
	return []Slave{
		Slave{Hostname: "mksuns31", Port: 1912},
		Slave{Hostname: "mksuns32", Port: 1912},
	}
}

func SlaveIndex(w http.ResponseWriter, r *http.Request) {
	slaves := getSlaves()
	json.NewEncoder(w).Encode(slaves)
}

func SlaveById(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["slaveId"]
	id64, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := uint(id64)
	for _,slave := range getSlaves() {
		if slave.Id == id {
			json.NewEncoder(w).Encode(slave)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	return
}