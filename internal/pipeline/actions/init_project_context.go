package actions

import (
	"os"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/project"
	"github.com/80LK/godev/internal/utils"

	"github.com/80lk/modlike"
	"golang.org/x/mod/modfile"
)

type InitProjectContext struct{}

func (a InitProjectContext) Plan(
	ctx *pipeline.Context,
) ([]pipeline.Patch, error) {
	goProjectFile := project.GetGoProjectFile(ctx.ProjectDir)
	exsist, err := utils.ExsistFile(goProjectFile)
	if err != nil {
		return nil, err
	}

	if exsist {
		data, err := os.ReadFile(goProjectFile)
		if err != nil {
			return nil, err
		}
		doc, err := modlike.Parse(data)
		if err != nil {
			return nil, err
		}

		err = project.ParseIn(doc, ctx.Project)
		if err != nil {
			return nil, err
		}
	}

	goModPath := project.GetGoModFile(ctx.ProjectDir)
	exsist, err = utils.ExsistFile(goModPath)
	if err != nil {
		return nil, err
	}

	if exsist {
		data, err := os.ReadFile(goModPath)
		if err != nil {
			return nil, err
		}

		ctx.Mod, err = modfile.Parse(goModPath, data, nil)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
