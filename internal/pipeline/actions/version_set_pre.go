package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type VersionSetPre struct {
	Value string
}

func (v VersionSetPre) Plan(ctx *context.Context) ([]patches.Patch, error) {
	ctx.GoProject.Project.Version.PreRelease = v.Value
	return nil, nil
}
