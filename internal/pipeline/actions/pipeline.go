package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type Pipeline struct {
	Actions []Action
}

func (p Pipeline) Plan(
	ctx *context.Context,
) ([]patches.Patch, error) {

	var patches []patches.Patch

	for _, action := range p.Actions {

		actionPatches, err := action.Plan(ctx)
		if err != nil {
			return nil, err
		}

		patches = append(
			patches,
			actionPatches...,
		)
	}

	return patches, nil
}

func (p *Pipeline) Add(actions ...Action) *Pipeline {
	p.Actions = append(p.Actions, actions...)
	return p
}
