package apiobjects

import (
	"net/http"
	"encoding/json"
)

type Slave struct {
	Id uint			`json:"id"`
	Hostname string		`json:"hostname"`
	Port uint		`json:"slave_port"`
}

func SlaveIndex(w http.ResponseWriter, r *http.Request) {
	slaves := []Slave{
		Slave{Id: 1, Hostname: "mksuns31", Port: 1912},
		Slave{Id: 2, Hostname: "mksuns32", Port: 1912},
	}
	json.NewEncoder(w).Encode(slaves)
}