package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"

	"github.com/80LK/godev/internal/pipeline"
)

type VersionSet struct {
	Value string
}

func (v VersionSet) Plan(ctx *context.Context) ([]pipeline.Patch, error) {
	context.Set(ctx, _OLD_VERSION_CONTEXT_KEY, ctx.GoProject.Project.Version.Clone())
	return nil, ctx.GoProject.Project.Version.EncodeFrom(v.Value)
}
