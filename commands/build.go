package commands

import (
	"github.com/80LK/godev/internal/pipeline"
	goAct "github.com/80LK/godev/internal/pipeline/actions/go"
	project "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/spf13/cobra"
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "short desc",
	Long:  "long desc",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.New(dryRun)

		return pipeline.New().Add(
			project.InitProjectContext{},
			goAct.Builds{},
		).Execute(ctx)
	},
}

func init() {
	Root.AddCommand(BuildCmd)
}
