package slaves

import (
	"net/http"
	"encoding/json"
)

type Slave struct {
	Hostname string
	Port uint
}

func SlaveIndex(w http.ResponseWriter, r *http.Request) {
	slaves := []Slave{
		Slave{Hostname: "mksuns31", Port: 1912},
		Slave{Hostname: "mksuns32", Port: 1912},
	}
	json.NewEncoder(w).Encode(slaves)
}