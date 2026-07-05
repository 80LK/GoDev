package actions

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/project"
)

type InitProject struct {
	ModuleName string
	Author     string
}

func (f InitProject) Plan(ctx *pipeline.Context) ([]pipeline.Patch, error) {
	if ctx.GoProject != nil {
		return nil, nil
	}

	ctx.GoProject = project.New()
	if err := project.ParseFromGoModuleNameIn(f.ModuleName, ctx.GoProject); err != nil {
		return nil, err
	}

	if ctx.GoProject.Project.Version.Major == 0 && ctx.GoProject.Project.Version.Minor == 0 && ctx.GoProject.Project.Version.Patch == 0 {
		ctx.GoProject.Project.Version.Minor = 1
	}

	if f.Author != "" {
		ctx.GoProject.Project.Author = f.Author
	}

	return nil, nil
}
