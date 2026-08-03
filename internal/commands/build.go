package commands

import (
	"github.com/80LK/godev/internal/pipeline"
	goAct "github.com/80LK/godev/internal/pipeline/actions/go"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/spf13/cobra"
)

var parallel bool
var BuildCmd = &cobra.Command{
	Use:   "build <...target>",
	Short: "build target",
	Long:  `Build target sources. If target not set, build all`,
	RunE: func(cmd *cobra.Command, targets []string) error {
		ctx := context.New(dryRun)

		return pipeline.New().Add(
			projectAct.InitProjectContext{},
			goAct.Builds{
				Targets:  targets,
				Parallel: parallel,
			},
		).Execute(ctx)
	},
}

func init() {
	BuildCmd.Flags().BoolVarP(&parallel, "parallel", "p", false, "paralleling builds")
	Root.AddCommand(BuildCmd)
}
