package apiobjects

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

type Slave struct {
	Id uint			   `json:"id"`
	Hostname string		   `json:"hostname"`
	Port uint		   `json:"slave_port"`
	MongodPortRangeBegin uint  `json:"mongod_port_range_begin"` //inclusive
	MongodPortRangeEnd uint    `json:"mongod_port_range_end"`   //exclusive
	PersistantStorage bool     `json:"persistant_storage"`
	RootDataDirectory string   `json:"root_data_directory"`
	State string               `json:"state"`
}

func getSlaves() []Slave {
	return []Slave{
		Slave{Id: 0, Hostname: "mksuns31", Port: 1912, MongodPortRangeBegin: 20000, MongodPortRangeEnd:20100, PersistantStorage:true, RootDataDirectory:"/home/mongo/data", State: "active"},
		Slave{Id: 1, Hostname: "mksuns32", Port: 1912, MongodPortRangeBegin: 20000, MongodPortRangeEnd:20001, PersistantStorage:false, RootDataDirectory:"/home/mongo/data", State: "active"},
		Slave{Id: 2, Hostname: "mksuns33", Port: 1912, MongodPortRangeBegin: 20000, MongodPortRangeEnd:20001, PersistantStorage:false, RootDataDirectory:"/home/mongo/data", State: "active"},
		Slave{Id: 3, Hostname: "mksuns34", Port: 1912, MongodPortRangeBegin: 20000, MongodPortRangeEnd:20001, PersistantStorage:false, RootDataDirectory:"/home/mongo/data", State: "active"},
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