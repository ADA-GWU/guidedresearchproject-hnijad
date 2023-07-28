package cmd

import (
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/config"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/sos"
	"github.com/spf13/cobra"
)

var dataNode = &cobra.Command{
	Use:     "data",
	Aliases: []string{"data"},
	Short:   "command to start data node",
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		volDir, _ := cmd.Flags().GetString("vol_dir")
		primaryNode, _ := cmd.Flags().GetString("primary_node")
		port, _ := cmd.Flags().GetString("port")
		grpcPort, _ := cmd.Flags().GetString("grpc_port")
		nodeId, _ := cmd.Flags().GetString("node_id")

		params := &config.DataNodeParams{
			NodeId:         nodeId,
			HttpPort:       port,
			GRPCPort:       grpcPort,
			VolDir:         volDir,
			PrimaryNodeUrl: primaryNode,
		}

		sos.RunDataNode(params)
	},
}

func init() {
	root.AddCommand(dataNode)
	dataNode.PersistentFlags().String("port", "8080", "Port to start the http server on")
	dataNode.PersistentFlags().String("grpc_port", "8080", "Grpc port to start the grpc server on")
	dataNode.PersistentFlags().String("vol_dir", "", "File system directory for volumes")
	dataNode.PersistentFlags().String("primary_node", "", "Primary node address")
	dataNode.PersistentFlags().String("node_id", "", "Data node id")
}
