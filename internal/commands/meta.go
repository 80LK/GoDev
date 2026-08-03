package commands

import (
	"github.com/80LK/godev/cli"
	"github.com/spf13/cobra"
)

var MetaCmd = &cobra.Command{
	Use:   "meta",
	Short: "Generate meta-file",
	Long:  "Generate meta-file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.GenerateMeta(cli.DryRunOptions{DryRun: dryRun})
	},
}

func init() {
	Root.AddCommand(MetaCmd)
}
