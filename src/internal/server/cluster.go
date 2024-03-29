package server

import (
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/client"
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ClusterInfo struct {
	Nodes               map[string]*pb.DataNodeInfo              `json:"nodes"`
	DataNodeGrpcClients map[string]*client.DataGrpcClientWrapper `json:"-"`
}

func (c *ClusterInfo) GetDataNodeGrpcClient(dataNodeId string, dataNodeUrl string) (*client.DataGrpcClientWrapper, error) {
	// TODO check dataNodeUrl argument
	if _, ok := c.DataNodeGrpcClients[dataNodeId]; ok {
		return c.DataNodeGrpcClients[dataNodeId], nil
	}
	c.DataNodeGrpcClients[dataNodeId] = client.NewDataGrpcClient(dataNodeUrl)
	return c.DataNodeGrpcClients[dataNodeId], nil
}

func (c *ClusterInfo) AddNewDataNode(info *pb.DataNodeInfo) error {
	info.LastHeartBeatAt = timestamppb.New(time.Now())
	if _, ok := c.Nodes[info.Id]; ok {
		log.Info("Updating existing data node info", info.Id)
		c.Nodes[info.Id] = info
	} else {
		log.Info("New data node is detected with the id", info.Id)
		c.Nodes[info.Id] = info
	}

	return nil
}

type VolumeResult struct {
	Id          int32  `json:"id"`
	DataNodeUrl string `json:"data_node_url"`
}

func (c *ClusterInfo) FindVolumeWithMaxAvailableSpace() *VolumeResult {
	log.Infoln("FindVolumeWithMaxAvailableSpace start")
	freeSpace := int64(-1)
	volumeID := int32(-1)
	dataNodeUrl := ""
	httpPort := ""

	for _, val := range c.Nodes {
		for _, volume := range val.Volumes {
			if volume.FreeSpace > freeSpace {
				freeSpace = volume.FreeSpace
				volumeID = volume.Id
				dataNodeUrl = val.Address
				httpPort = val.HttpPort
			}
		}
	}
	log.Infoln("FindVolumeWithMaxAvailableSpace end")

	return &VolumeResult{
		DataNodeUrl: fmt.Sprintf("%v:%v", dataNodeUrl, httpPort),
		Id:          volumeID,
	}
}

func (c *ClusterInfo) findDataNodeWithVolumeID(volumeId int32) *pb.DataNodeInfo {
	for _, val := range c.Nodes {
		for _, volume := range val.Volumes {
			if volume.Id == volumeId {
				return val
			}
		}
	}
	return nil
}
