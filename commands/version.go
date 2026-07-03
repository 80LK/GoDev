package commands

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "short desc",
	Long:  "long desc",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	Root.AddCommand(versionCmd)
}
