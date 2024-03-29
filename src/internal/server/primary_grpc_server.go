package server

import (
	"context"
	"fmt"
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"time"
)

func (s *PrimaryServer) HeartBeat(context context.Context, request *pb.DataNodeInfo) (*emptypb.Empty, error) {
	//log.Info("Got heartbeat from the node", request.Id)
	_ = s.ClusterInfo.AddNewDataNode(request)
	return &emptypb.Empty{}, nil
}

func (s *PrimaryServer) VolumeWatcher() {

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Starting volume checking...")
			s.manage()
			log.Println("Finished volume checking")

		}
	}
}

func (s *PrimaryServer) manage() {
	// Rules for creating new volumes
	for key, val := range s.ClusterInfo.Nodes {
		//log.Info("key = ", key, "val = ", val)
		if len(val.Volumes) == 0 { // Rule 1, if no volume exists on data node, create one
			volumeId, _ := s.State.NextVolumeId()
			client, _ := s.ClusterInfo.GetDataNodeGrpcClient(key, val.Address+":"+val.GrpcPort) // TODO change address to grpc
			_, err := client.CreateVolume(&pb.VolumeCreateRequest{VolumeId: int32(volumeId)})   // TODO change types to avoid type conversions
			if err != nil {
				log.Errorf("Attempt to create new volume failed on %v with error %v", key, err.Error())
			}
		}
	}
}

func StartPrimaryNodeGrpcServer(primaryServer *PrimaryServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", primaryServer.Params.GRPCPort))
	if err != nil {
		log.Info("error happened", err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterPrimaryNodeServer(s, primaryServer) // add the same primary node reference
	log.Infoln("Starting the grpc server at", primaryServer.Params.GRPCPort)
	err = s.Serve(lis)
	if err != nil {
		log.Infoln("error happened", err.Error())
		return
	}
}
