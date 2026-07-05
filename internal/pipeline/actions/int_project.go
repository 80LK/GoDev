package actions

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/project"
)

type IntProject struct{}

func (i IntProject) Plan(ctx *pipeline.Context) ([]pipeline.Patch, error) {

	if ctx.Project == nil {
		ctx.Project = project.New()
	}

	if err := project.ParseFromGoModIn(ctx.Mod, ctx.Project); err != nil {
		return nil, err
	}

	if ctx.Project.Project.Version.Major == 0 && ctx.Project.Project.Version.Minor == 0 && ctx.Project.Project.Version.Patch == 0 {
		ctx.Project.Project.Version.Minor = 1
	}
	return nil, nil
}
