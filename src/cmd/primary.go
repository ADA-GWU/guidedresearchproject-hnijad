package cmd

import (
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/config"
	"github.com/ADA-GWU/guidedresearchproject-hnijad/internal/sos"
	"github.com/spf13/cobra"
)

var primaryNode = &cobra.Command{
	Use:     "primary",
	Aliases: []string{"primary"},
	Short:   "command to start primary node",
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		port, _ := cmd.Flags().GetString("port")
		grpcPort, _ := cmd.Flags().GetString("grpc_port")

		params := &config.PrimaryNodeParams{
			HttpPort: port,
			GRPCPort: grpcPort,
		}

		sos.RunPrimaryNode(params)
	},
}

func init() {
	root.AddCommand(primaryNode)
	primaryNode.PersistentFlags().String("port", "8080", "Port to start the http server on")
	primaryNode.PersistentFlags().String("grpc_port", "8080", "Grpc port to start the grpc server on")
}
