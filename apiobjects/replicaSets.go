package apiobjects

import (
	"net/http"
	"encoding/json"
)

type ReplicaSet struct {
	Id uint					`json:"id"`
	Name string				`json:"name"`
	PersistentNodeCunt uint			`json:"presistent_node_count"`
	VolatileNodeCount uint			`json:"volatile_node_count"`
	ConfigureAsShardingConfigServer bool	`json:"configure_as_sharding_config_server"`
}

func ReplicaSetIndex(w http.ResponseWriter, r *http.Request) {
	replicaSets := []ReplicaSet{
		ReplicaSet{Id: 1, Name: "meterologic_data", PersistentNodeCunt: 1, VolatileNodeCount: 2, ConfigureAsShardingConfigServer: false},
	}
	json.NewEncoder(w).Encode(replicaSets)
}