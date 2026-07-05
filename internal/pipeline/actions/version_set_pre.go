package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"

	"github.com/80LK/godev/internal/pipeline"
)

type VersionSetPre struct {
	Value string
}

func (v VersionSetPre) Plan(ctx *context.Context) ([]pipeline.Patch, error) {
	ctx.GoProject.Project.Version.PreRelease = v.Value
	return nil, nil
}
