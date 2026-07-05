package actions

import (
	"github.com/80LK/godev/internal/pipeline"
)

type VersionSetPre struct {
	Value string
}

func (v VersionSetPre) Plan(ctx *pipeline.Context) ([]pipeline.Patch, error) {
	ctx.Project.Project.Version.SetPreRelease(v.Value)
	return nil, nil
}
