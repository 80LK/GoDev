package cli

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/actions"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/project"

	fsAct "github.com/80LK/godev/internal/pipeline/actions/fs"
	gitAct "github.com/80LK/godev/internal/pipeline/actions/git"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
)

func Release(opts DryRunOptions) error {
	ctx := context.New(opts.DryRun)

	return pipeline.New().Add(
		actions.OR{
			Actions: []actions.Action{
				fsAct.CheckExistsFile{Path: project.GetJSONProjectFile(ctx.ProjectDir)},
				fsAct.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			},
		},
		gitAct.CheckClearGit{},
		projectAct.InitProjectContext{},
		projectAct.RunScript{IgnoreNotFound: true, Name: project.LifecycleName(project.PhaseBefore, "release")},
		pipeline.New().Add(
			projectAct.GenerateMeta{},
			gitAct.CheckClearGit{OutputKey: "git_clear"},
			actions.ConditionAction{
				Condition: func(ctx *context.Context) bool {
					res, _ := context.Get[bool](ctx, "git_clear")
					return !res
				},
				Action: gitAct.GitCommit{Value: "generate meta info"},
			},
		),
		gitAct.GitTagVersion{},
		projectAct.RunScript{IgnoreNotFound: true, Name: project.LifecycleName(project.PhaseBefore, "release")},
	).Execute(ctx)
}
