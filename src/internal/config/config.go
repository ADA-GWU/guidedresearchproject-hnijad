package config

type DataNodeParams struct {
	NodeId         string
	HttpPort       string
	GRPCPort       string
	VolDir         string
	PrimaryNodeUrl string
}

type PrimaryNodeParams struct {
	HttpPort string
	GRPCPort string
}
