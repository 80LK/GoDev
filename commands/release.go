package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "short desc",
	Long:  "long desc",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("god release not implemented")
	},
}

func init() {
	Root.AddCommand(releaseCmd)
}
