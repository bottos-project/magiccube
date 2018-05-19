package util

type NodeDBInfo struct {
	NodeId      string   `json:"node_id"`
	NodeIP      string   `json:"node_ip"`
	NodePort    string   `json:"node_port"`
	NodeAddress string   `json:"node_address"`
	SeedIP      string   `json:"seed_ip"`
	SlaveIP     []string `json:"slave_ip"`
}
