package server

import (
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/config"
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
)

type PrimaryServer struct {
	pb.UnimplementedPrimaryNodeServer
	Params      *config.PrimaryNodeParams
	ClusterInfo *ClusterInfo
}

func (s *PrimaryServer) GetClusterInfo() *ClusterInfo {
	return s.ClusterInfo
}

func NewPrimaryServer(params *config.PrimaryNodeParams, info *ClusterInfo) *PrimaryServer {
	return &PrimaryServer{
		Params:      params,
		ClusterInfo: info,
	}
}
