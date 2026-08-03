package cli

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/context"

	goAct "github.com/80LK/godev/internal/pipeline/actions/go"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
)

type BuildOptions struct {
	DryRunOptions

	Targets  []string
	Parallel bool
}

func Builds(opts BuildOptions) error {
	ctx := context.New(opts.DryRun)

	return pipeline.New().Add(
		projectAct.InitProjectContext{},
		goAct.Builds{
			Targets:  opts.Targets,
			Parallel: opts.Parallel,
		},
	).Execute(ctx)
}
