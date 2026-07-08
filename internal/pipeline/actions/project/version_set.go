package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type VersionSet struct {
	OldVersionKey string
	NewVersionKey string
	Value         string
}

func (v VersionSet) Plan(ctx *context.Context) ([]patches.Patch, error) {
	if v.OldVersionKey != "" {
		context.Set(ctx, v.OldVersionKey, ctx.GoProject.Project.Version.Clone())
	}

	if err := ctx.GoProject.Project.Version.EncodeFrom(v.Value); err != nil {
		return nil, err
	}

	if v.NewVersionKey != "" {
		context.Set(ctx, v.NewVersionKey, ctx.GoProject.Project.Version.Clone())
	}

	return nil, nil
}
