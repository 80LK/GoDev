package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/project"
)

type IntProject struct {
	Author string
}

func (i IntProject) Plan(ctx *context.Context) ([]pipeline.Patch, error) {

	if ctx.GoProject == nil {
		ctx.GoProject = project.New()
	}

	if err := project.ParseFromGoModIn(ctx.Mod, ctx.GoProject); err != nil {
		return nil, err
	}

	if ctx.GoProject.Project.Version.Major == 0 && ctx.GoProject.Project.Version.Minor == 0 && ctx.GoProject.Project.Version.Patch == 0 {
		ctx.GoProject.Project.Version.Minor = 1
	}

	if i.Author != "" {
		ctx.GoProject.Project.Author = i.Author
	}
	return nil, nil
}
