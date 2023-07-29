package server

import (
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/config"
	pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"
)

type PrimaryServer struct {
	pb.UnimplementedPrimaryNodeServer
	Params *config.PrimaryNodeParams
}

func NewPrimaryServer(params *config.PrimaryNodeParams) *PrimaryServer {
	return &PrimaryServer{
		Params: params,
	}
}
