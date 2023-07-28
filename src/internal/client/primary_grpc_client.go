package client

import (
	"context"
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type PrimaryGrpcClientWrapper struct {
	masterGrpcClient pb.PrimaryNodeClient
}

func NewMasterGrpcClient(primaryUrl string) *PrimaryGrpcClientWrapper {
	conn, err := grpc.Dial(primaryUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}
	c := pb.NewPrimaryNodeClient(conn)
	return &PrimaryGrpcClientWrapper{
		masterGrpcClient: c,
	}
}

func (s *PrimaryGrpcClientWrapper) HeartBeat(message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := s.masterGrpcClient.HeartBeat(ctx,
		&pb.DataNodeInfo{
			Id: message,
		},
	)
	if err != nil {
		log.Fatalf("could not heartbeat: %v", err)
	}
}
