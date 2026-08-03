package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type OR struct {
	Actions []Action
}

// Plan implements [Action].
func (o OR) Plan(ctx *context.Context) ([]patches.Patch, error) {
	var gError error
	for _, action := range o.Actions {
		patches, err := action.Plan(ctx)
		if err != nil {
			if gError == nil {
				gError = err
			}
			continue
		}

		return patches, nil
	}
	return nil, gError
}
