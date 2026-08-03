package commands

import (
	"github.com/80LK/godev/cli"
	"github.com/spf13/cobra"
)

var ReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Set git tag for release",
	Long:  "Set git tag for release",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.Release(cli.DryRunOptions{DryRun: dryRun})
	},
}

func init() {
	Root.AddCommand(ReleaseCmd)
}
