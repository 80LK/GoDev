package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type Execute struct {
	Execute func(ctx *context.Context)
}

func (e Execute) Plan(ctx *context.Context) ([]patches.Patch, error) {
	e.Execute(ctx)
	return nil, nil
}
