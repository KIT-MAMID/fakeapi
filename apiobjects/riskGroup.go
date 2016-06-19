package apiobjects

import (
	"net/http"
	"encoding/json"
)

type RiskGroup struct {
	Id uint					`json:"id"`
	Name string				`json:"name"`
}

func RiskGroupIndex(w http.ResponseWriter, r *http.Request) {
	replicaSets := []ReplicaSet{
		ReplicaSet{Id: 1, Name: "Rack A"},
		ReplicaSet{Id: 2, Name: "Rack B"},
	}
	json.NewEncoder(w).Encode(replicaSets)
}