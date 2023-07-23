package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var root = &cobra.Command{
	Use:   "sos",
	Short: "sos - a simple object storage",
	Long:  "sos - a simple object storage optimized for small objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := root.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing command '%s'", err)
		os.Exit(1)
	}
}
