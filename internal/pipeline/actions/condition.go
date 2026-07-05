package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"

	"github.com/80LK/godev/internal/pipeline"
)

type ConditionAction struct {
	Condition func(
		ctx *context.Context,
	) bool
	Action pipeline.Action
}

func (a ConditionAction) Plan(
	ctx *context.Context,
) ([]pipeline.Patch, error) {
	if a.Condition(ctx) {
		return a.Action.Plan(ctx)
	}

	return nil, nil
}
