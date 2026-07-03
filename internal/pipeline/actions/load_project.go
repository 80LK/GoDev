package actions

import (
	"os"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/project"
	"github.com/80LK/godev/internal/utils"

	"github.com/80lk/modlike"
	"golang.org/x/mod/modfile"
)

type LoadProject struct{}

func (a LoadProject) Plan(
	ctx *pipeline.Context,
) ([]pipeline.Patch, error) {

	goProjectFile := project.GetGoProjectFile(ctx.ProjectDir)
	exsist, err := utils.ExsistFile(goProjectFile)
	if err != nil {
		return nil, err
	}

	if exsist {
		ctx.HasGoProject = true
		data, err := os.ReadFile(goProjectFile)
		if err != nil {
			return nil, err
		}
		doc, err := modlike.Parse(data)
		if err != nil {
			return nil, err
		}

		ctx.Project, err = project.Parse(doc)
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
		ctx.HasGoMod = true

		if !ctx.HasGoProject {
			data, err := os.ReadFile(goModPath)
			if err != nil {
				return nil, err
			}

			file, err := modfile.Parse(goModPath, data, nil)
			if err != nil {
				return nil, err
			}

			ctx.Project, err = project.ParseFromGoMod(file)
			if err != nil {
				return nil, err
			}
		}

	}

	if ctx.Project == nil {
		ctx.Project = &project.GoProject{
			Project: &project.ProjectInfo{},
		}
	}

	if ctx.Project.Project.Name == "" {
	}

	return nil, nil
}
