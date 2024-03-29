package server

import (
	"fmt"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/client"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/config"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/storage"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
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

	for i := 0; i < 8192/2; i++ {
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

func TestDataServer_WriteObject(t *testing.T) {
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
	readPath := "/Users/hnijad/Desktop/lab/read-test/"
	i := 0

	dirEntries, err := os.ReadDir(readPath)

	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}

		fileName := dirEntry.Name()

		filePath := filepath.Join(readPath, fileName)

		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			continue
		}
		fid := fmt.Sprintf("16,%v", i)
		err = ds.WriteObject(fid, fileName, nil, data)
		if err != nil {
			log.Fatalf("Error when writing object")
		}
		i++
	}
}
