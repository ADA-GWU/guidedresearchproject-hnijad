package server

import (
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/api"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/config"
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
)

type PrimaryServer struct {
	pb.UnimplementedPrimaryNodeServer
	Params      *config.PrimaryNodeParams
	ClusterInfo *ClusterInfo
	State       *PrimaryNodeState
}

func (s *PrimaryServer) GetClusterInfo() *ClusterInfo {
	return s.ClusterInfo
}

func (s *PrimaryServer) FindDataNode() (*api.VolumeRequest, error) {
	volumeResult := s.ClusterInfo.FindVolumeWithMaxAvailableSpace()
	objectId, _ := s.State.NextId()
	return &api.VolumeRequest{
		ObjectId:    fmt.Sprintf("%v,%v", volumeResult.Id, objectId),
		DataNodeUrl: volumeResult.DataNodeUrl,
	}, nil
}

func NewPrimaryServer(params *config.PrimaryNodeParams, info *ClusterInfo, state *PrimaryNodeState) *PrimaryServer {
	return &PrimaryServer{
		Params:      params,
		ClusterInfo: info,
		State:       state,
	}
}
