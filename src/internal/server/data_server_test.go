package server

import (
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/client"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/config"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/storage"
	"testing"
)

func TestDataServer_ReadObject(t *testing.T) {

	params := &config.DataNodeParams{
		NodeId:         "1",
		HttpPort:       "8081",
		GRPCPort:       "1235",
		VolDir:         "/Users/hnijad/Desktop/lab/guidedresearchproject-hnijad/src/tmp/node1",
		PrimaryNodeUrl: "localhost:1234",
	}
	dataStorage := storage.NewStorage(params.VolDir)

	ds := &DataServer{
		ID:                params.NodeId,
		Storage:           dataStorage,
		PrimaryGrpcClient: client.NewMasterGrpcClient(params.PrimaryNodeUrl),
		Params:            params,
	}

	for i := 0; i < 100000; i++ {
		_, err := ds.ReadObject(fmt.Sprintf("16,%v", i))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkDataServer_ReadObject(b *testing.B) {
	params := &config.DataNodeParams{
		NodeId:         "1",
		HttpPort:       "8081",
		GRPCPort:       "1235",
		VolDir:         "/Users/hnijad/Desktop/lab/guidedresearchproject-hnijad/src/tmp/node1",
		PrimaryNodeUrl: "localhost:1234",
	}
	dataStorage := storage.NewStorage(params.VolDir)

	ds := &DataServer{
		ID:                params.NodeId,
		Storage:           dataStorage,
		PrimaryGrpcClient: client.NewMasterGrpcClient(params.PrimaryNodeUrl),
		Params:            params,
	}

	for i := 0; i < b.N; i++ {
		_, err := ds.ReadObject(fmt.Sprintf("16,%v", 1))
		if err != nil {
			fmt.Println(err)
		}
	}
}
