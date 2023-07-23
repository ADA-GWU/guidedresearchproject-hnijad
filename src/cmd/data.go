package cmd

import (
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
		nodeId, _ := cmd.Flags().GetString("node_id")
		sos.RunDataNode(volDir, port, primaryNode, nodeId)
	},
}

func init() {
	root.AddCommand(dataNode)
	dataNode.PersistentFlags().String("port", "8080", "Port to start the http server on")
	dataNode.PersistentFlags().String("vol_dir", "", "File system directory for volumes")
	dataNode.PersistentFlags().String("primary_node", "", "Primary node address")
	dataNode.PersistentFlags().String("node_id", "", "Datan node id")
}
