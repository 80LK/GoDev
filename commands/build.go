package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "short desc",
	Long:  "long desc",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("god build not implemented")
	},
}

func init() {
	Root.AddCommand(buildCmd)
}
