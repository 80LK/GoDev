package commands

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/project"
	"github.com/spf13/cobra"

	fsAct "github.com/80LK/godev/internal/pipeline/actions/fs"
	gitAct "github.com/80LK/godev/internal/pipeline/actions/git"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate from old tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.New(dryRun)

		return pipeline.New().Add(
			fsAct.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			gitAct.CheckClearGit{},
			projectAct.InitProjectContext{DisabelDeprecetedMessage: true},
			projectAct.Migrate{},
			gitAct.GitCommit{Value: "migrate project structure"},
		).Execute(ctx)
	},
}

func init() {
	Root.AddCommand(MigrateCmd)
}
