package models

type StatMessage struct {
	NodeId string `json:"node_id"`
	CommunityChain string `json:"community_chain"`
	TotalMemory int64 `json:"total_memory"`
	UsedMemory int64 `json:"used_memory"`
	Processors []float64 `json:"processors"`
	NetInterfaces []NetIntarfaceStats `json:"net_interfaces"`
}

type NetIntarfaceStats struct {
	InterfaceName string `json:"interface_name"`
	IncomeBytes int64 `json:"income_bytes"`
	OutComeBytes int64 `json:"outcome_bytes"`
}
