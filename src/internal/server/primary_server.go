package server

import pb "github.com/ADA-GWU/guidedresearchproject-hnijad/internal/proto/primary"

type PrimaryServer struct {
	pb.UnimplementedPrimaryNodeServer
}

func NewPrimaryServer() *PrimaryServer {
	return &PrimaryServer{}
}
