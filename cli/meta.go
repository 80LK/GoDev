package cli

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/actions"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/project"

	fsAct "github.com/80LK/godev/internal/pipeline/actions/fs"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
)

func GenerateMeta(opts DryRunOptions) error {
	ctx := context.New(opts.DryRun)

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
}
