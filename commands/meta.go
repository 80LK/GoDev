package commands

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/actions"
	fsAct "github.com/80LK/godev/internal/pipeline/actions/fs"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/project"
	"github.com/spf13/cobra"
)

var MetaCmd = &cobra.Command{
	Use:   "meta",
	Short: "Generate meta-file",
	Long:  "Generate meta-file",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.New(dryRun)

		return pipeline.New().Add(
			fsAct.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			projectAct.InitProjectContext{},
			actions.ConditionAction{
				Condition: func(ctx *context.Context) bool {
					return ctx.GoProject.Meta
				},
				Action: projectAct.GenerateMeta{},
			},
		).Execute(ctx)
	},
}

func init() {
	Root.AddCommand(MetaCmd)
}
