package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"

	"github.com/80LK/godev/internal/pipeline"
)

type VersionSet struct {
	Value string
}

func (v VersionSet) Plan(ctx *context.Context) ([]pipeline.Patch, error) {
	oldVersion := ctx.GoProject.Project.Version.Clone()
	if err := ctx.GoProject.Project.Version.EncodeFrom(v.Value); err != nil {
		return nil, err
	}

	return (PatchSources{OldVersion: oldVersion}).Plan(ctx)
}
