package commands

import (
	"github.com/80LK/godev/internal/pipeline"
	goAct "github.com/80LK/godev/internal/pipeline/actions/go"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/spf13/cobra"
)

var BuildCmd = &cobra.Command{
	Use:   "build <target>",
	Short: "build target",
	Long:  `Build target sources. If target not set, build all`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.New(dryRun)
		target := ""

		if len(args) == 1 {
			target = args[0]
		}

		return pipeline.New().Add(
			projectAct.InitProjectContext{},
			goAct.Builds{Target: target},
		).Execute(ctx)
	},
}

func init() {
	Root.AddCommand(BuildCmd)
}
