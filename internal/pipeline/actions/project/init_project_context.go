package actions

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"

	"github.com/80LK/godev/internal/utils"
	"github.com/80LK/godev/project"

	"github.com/80LK/modlike"
	"golang.org/x/mod/modfile"
)

type InitProjectContext struct {
	DisabelDeprecetedMessage bool
}

func (a InitProjectContext) Plan(
	ctx *context.Context,
) ([]patches.Patch, error) {
	goProjectFile := project.GetGoProjectFile(ctx.ProjectDir)
	exsist, err := utils.ExsistFile(goProjectFile)
	if err != nil {
		return nil, err
	}

	if exsist {
		if !a.DisabelDeprecetedMessage {
			fmt.Printf("go.project depreceted file. use \"god migrate\" for migrate new config file. Support modlike was been remove in god version 0.2")
		}
		ctx.DeprecetedConfig = true

		data, err := os.ReadFile(goProjectFile)
		if err != nil {
			return nil, err
		}

		ctx.GoProject = new(project.GoProject)

		err = modlike.Unmarshal(data, ctx.GoProject)
		if err != nil {
			return nil, err
		}
	} else {
		jsonProjectFile := project.GetJSONProjectFile(ctx.ProjectDir)
		exsist, err = utils.ExsistFile(goProjectFile)
		if err != nil {
			return nil, err
		}

		if exsist {
			data, err := os.ReadFile(jsonProjectFile)
			if err != nil {
				return nil, err
			}

			ctx.GoProject = new(project.GoProject)

			err = json.Unmarshal(data, ctx.GoProject)
			if err != nil {
				return nil, err
			}
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
