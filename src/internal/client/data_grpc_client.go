package client

import (
	"context"
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type DataGrpcClientWrapper struct {
	dataGrpcClient pb.DataNodeClient
}

func NewDataGrpcClient(dataNodeUrl string) *DataGrpcClientWrapper {
	conn, err := grpc.Dial(dataNodeUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}
	c := pb.NewDataNodeClient(conn)
	return &DataGrpcClientWrapper{
		dataGrpcClient: c,
	}
}

func (c *DataGrpcClientWrapper) CreateVolume(request *pb.VolumeCreateRequest) (*pb.VolumeCreateResponse, error) {
	//TODO check method signature
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := c.dataGrpcClient.CreateVolume(ctx, request)
	return resp, err
}
