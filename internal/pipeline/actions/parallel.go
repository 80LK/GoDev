package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/utils"
)

type Parallel struct {
	Items []Action
}

func (a Parallel) Plan(
	ctx *context.Context,
) ([]patches.Patch, error) {

	items, err := utils.Map(a.Items, func(item Action, _ int, _ []Action) ([]patches.Patch, error) {
		return item.Plan(ctx)
	})
	if err != nil {
		return nil, err
	}

	return []patches.Patch{
		patches.ParallelPatch{
			Items: items,
		},
	}, nil
}
