package actions

import (
	"maps"
	"slices"

	"github.com/80LK/godev/internal/pipeline/actions"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type Builds struct {
	Targets  []string
	Parallel bool
}

func (b Builds) Plan(ctx *context.Context) ([]patches.Patch, error) {
	if ctx.GoProject.Builds == nil {
		return nil, nil
	}

	var plan actions.Executor

	if b.Parallel {
		plan = &actions.Parallel{}
	} else {
		plan = &actions.Pipeline{}
	}

	if b.Targets == nil {
		b.Targets = slices.Collect(maps.Keys(ctx.GoProject.Builds))
	}

	for _, buildName := range b.Targets {
		plan.Add(Build{Target: buildName})
	}

	return plan.Plan(ctx)
}
