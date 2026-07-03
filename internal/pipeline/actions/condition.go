package actions

import (
	"github.com/80LK/godev/internal/pipeline"
)

type ConditionAction struct {
	Condition func(
		ctx *pipeline.Context,
	) bool
	Action pipeline.Action
}

func (a ConditionAction) Plan(
	ctx *pipeline.Context,
) ([]pipeline.Patch, error) {
	if a.Condition(ctx) {
		return a.Action.Plan(ctx)
	}

	return nil, nil
}
