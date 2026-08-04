package actions

import (
	"maps"
	"slices"

	"github.com/80LK/godev/internal/pipeline/actions"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/utils/logger"
)

type Builds struct {
	Targets  []string
	Parallel bool
}

func (b Builds) Plan(ctx *context.Context) ([]patches.Patch, error) {
	logger := logger.Get("BUILD")
	if ctx.GoProject.Builds == nil {
		return nil, nil
	}

	var plan actions.Executor

	logger.Log("Parallel: %t", b.Parallel)
	if b.Parallel {
		plan = &actions.Parallel{}
	} else {
		plan = &actions.Pipeline{}
	}

	logger.Log("Targets len: %d", len(b.Targets))
	if len(b.Targets) == 0 {
		b.Targets = slices.Collect(maps.Keys(ctx.GoProject.Builds))
	}

	for _, buildName := range b.Targets {
		plan.Add(Build{Target: buildName})
	}

	return plan.Plan(ctx)
}
