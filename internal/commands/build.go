package commands

import (
	"github.com/80LK/godev/cli"
	"github.com/spf13/cobra"
)

var parallel bool
var BuildCmd = &cobra.Command{
	Use:   "build <...target>",
	Short: "build target",
	Long:  `Build target sources. If target not set, build all`,
	RunE: func(cmd *cobra.Command, targets []string) error {
		opts := cli.BuildOptions{}
		opts.DryRun = dryRun
		opts.Parallel = parallel
		opts.Targets = targets

		return cli.Builds(opts)
	},
}

func init() {
	BuildCmd.Flags().BoolVarP(&parallel, "parallel", "p", false, "paralleling builds")
	Root.AddCommand(BuildCmd)
}
