package server

import (
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ClusterInfo struct {
	Nodes map[string]*pb.DataNodeInfo `json:"nodes"`
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
