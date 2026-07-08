package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type VersionSetPre struct {
	OldVersionKey string
	Value         string
}

func (v VersionSetPre) Plan(ctx *context.Context) ([]patches.Patch, error) {
	if v.OldVersionKey != "" {
		context.Set(ctx, v.OldVersionKey, ctx.GoProject.Project.Version.Clone())
	}
	ctx.GoProject.Project.Version.PreRelease = v.Value
	return nil, nil
}
