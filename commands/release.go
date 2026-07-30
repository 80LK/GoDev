package commands

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/actions"
	fsAct "github.com/80LK/godev/internal/pipeline/actions/fs"
	gitAct "github.com/80LK/godev/internal/pipeline/actions/git"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/project"
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
			gitAct.CheckClearGit{},
			projectAct.InitProjectContext{},
			projectAct.RunScript{IgnoreNotFound: true, Name: project.LifecycleName(project.PhaseBefore, "release")},
			actions.ConditionAction{
				Condition: func(ctx *context.Context) bool {
					return ctx.GoProject.Meta
				},
				Action: pipeline.New().Add(
					projectAct.GenerateMeta{
						ContextKey: "generated",
					},
					actions.ConditionAction{
						Condition: func(ctx *context.Context) bool {
							res, _ := context.Get[bool](ctx, "generated")
							return res
						},
						Action: gitAct.GitCommit{Value: "generate meta info"},
					},
				),
			},
			gitAct.GitTagVersion{},
			projectAct.RunScript{IgnoreNotFound: true, Name: project.LifecycleName(project.PhaseBefore, "release")},
		).Execute(ctx)
	},
}

func init() {
	Root.AddCommand(ReleaseCmd)
}
