package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"

	"github.com/80LK/godev/internal/pipeline"
)

type VersionSetPre struct {
	Value string
}

func (v VersionSetPre) Plan(ctx *context.Context) ([]pipeline.Patch, error) {
	context.Set(ctx, _OLD_VERSION_CONTEXT_KEY, ctx.GoProject.Project.Version.Clone())
	ctx.GoProject.Project.Version.PreRelease = v.Value
	return nil, nil
}
