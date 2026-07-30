package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"

	"github.com/80LK/godev/project"
)

type WriteInProject struct {
	ModuleName string
	Author     string
}

func (f WriteInProject) Plan(ctx *context.Context) ([]patches.Patch, error) {
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
