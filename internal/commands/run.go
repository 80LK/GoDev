package commands

import (
	"fmt"

	"github.com/80LK/godev/cli"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run <script>",
	Short: "Run script",
	Long:  "Run script by name <script>",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Expect only name script")
		}

		opts := cli.RunScriptOptions{}
		opts.DryRun = dryRun
		opts.Name = args[0]

		return cli.RunScript(opts)
	},
}

func init() {
	Root.AddCommand(RunCmd)
}
