package cli

import (
	"github.com/80LK/godev/internal/pipeline"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
)

type RunScriptOptions struct {
	DryRunOptions

	Name string
}

func RunScript(opts RunScriptOptions) error {
	ctx := context.New(opts.DryRun)

	return pipeline.New().Add(
		projectAct.InitProjectContext{},
		projectAct.RunScript{
			Name: opts.Name,
		},
	).Execute(ctx)
}
