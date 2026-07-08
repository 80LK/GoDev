package commands

import (
	"github.com/80LK/godev/internal/pipeline"
	fsAct "github.com/80LK/godev/internal/pipeline/actions/fs"
	gitAct "github.com/80LK/godev/internal/pipeline/actions/git"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/project"
	"github.com/spf13/cobra"
)

var ReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Set git tag for release",
	Long:  "Set git tag for release",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.New(dryRun)

		return pipeline.New().Add(
			fsAct.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			projectAct.InitProjectContext{},

			gitAct.CheckClearGit{},
			gitAct.GitTagVersion{},
		).Execute(ctx)
	},
}

func init() {
	Root.AddCommand(ReleaseCmd)
}
