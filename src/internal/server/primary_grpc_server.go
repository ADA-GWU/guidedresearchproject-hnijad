package server

import (
	"context"
	"fmt"
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

func (s *PrimaryServer) HeartBeat(context context.Context, request *pb.DataNodeInfo) (*emptypb.Empty, error) {
	log.Info("Got heartbeat from the node", request.Id)
	return &emptypb.Empty{}, nil
}

func StartPrimaryNodeGrpcServer(grpcPort string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
	if err != nil {
		log.Info("error happened", err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterPrimaryNodeServer(s, &PrimaryServer{}) // add the same primary node reference
	log.Infoln("Starting the grpc server at", grpcPort)
	err = s.Serve(lis)
	if err != nil {
		log.Infoln("error happened", err.Error())
		return
	}
}
