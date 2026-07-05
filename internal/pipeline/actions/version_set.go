package actions

import (
	"github.com/80LK/godev/internal/pipeline"
)

type VersionSet struct {
	Value string
}

func (v VersionSet) Plan(ctx *pipeline.Context) ([]pipeline.Patch, error) {
	return nil, ctx.GoProject.Project.Version.EncodeFrom(v.Value)
}
