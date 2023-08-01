package server

import (
	"errors"
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/api"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/config"
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
	"strconv"
	"strings"
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
func (s *PrimaryServer) FindDataNodeByObjectId(objectId string) (*api.VolumeRequest, error) {
	tokens := strings.Split(objectId, ",")
	volumeId, err := strconv.ParseInt(tokens[0], 10, 32)

	if err != nil {
		return nil, errors.New("wrong id format")
	}

	res := s.ClusterInfo.findDataNodeWithVolumeID(int32(volumeId))

	if res == nil {
		return nil, errors.New("not Found")
	}

	return &api.VolumeRequest{
		ObjectId:    objectId,
		DataNodeUrl: fmt.Sprintf("%v:%v", res.Address, res.HttpPort),
	}, nil
}

func NewPrimaryServer(params *config.PrimaryNodeParams, info *ClusterInfo, state *PrimaryNodeState) *PrimaryServer {
	server := &PrimaryServer{
		Params:      params,
		ClusterInfo: info,
		State:       state,
	}
	go server.VolumeWatcher()
	return server
}
