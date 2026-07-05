package commands

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/actions"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/project"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Set git tag for release",
	Long:  "Set git tag for release",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.New(dryRun)

		return pipeline.New().Add(
			actions.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			actions.InitProjectContext{},

			actions.CheckClearGit{},
			actions.GitTagVersion{},
		).Execute(ctx)
	},
}

func init() {
	Root.AddCommand(releaseCmd)
}
