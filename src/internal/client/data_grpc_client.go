package client

import (
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
