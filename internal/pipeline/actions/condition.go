package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type ConditionAction struct {
	Condition func(
		ctx *context.Context,
	) bool
	Action Action
}

func (a ConditionAction) Plan(
	ctx *context.Context,
) ([]patches.Patch, error) {
	if a.Condition(ctx) {
		return a.Action.Plan(ctx)
	}

	return nil, nil
}
