package server

import (
	"context"
	"fmt"
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func (ds *DataServer) CreateVolume(ctx context.Context, req *pb.VolumeCreateRequest) (*pb.VolumeCreateResponse, error) {
	log.Infoln("starting to create a volume with id = ", req.VolumeId)
	if err := ds.Storage.CreateNewVolume(int(req.VolumeId)); err != nil {
		return nil, err
	}
	return &pb.VolumeCreateResponse{
		Status:  "success",
		Message: "successfully created",
	}, nil
}

func StartDataNodeGrpcServer(dataServer *DataServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", dataServer.Params.GRPCPort))
	if err != nil {
		log.Info("error happened", err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterDataNodeServer(s, dataServer) // add the same primary node reference
	log.Infoln("Starting the grpc dataServer at", dataServer.Params.GRPCPort)
	err = s.Serve(lis)
	if err != nil {
		log.Infoln("error happened ", err.Error())
		return
	}
}
