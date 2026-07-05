package actions

import (
	"fmt"
	"os"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/project"
)

type EncodeGoMod struct {
}

func (a EncodeGoMod) Plan(
	ctx *pipeline.Context,
) ([]pipeline.Patch, error) {

	if ctx.Mod == nil {
		return nil,
			fmt.Errorf(
				"project not found in context",
			)
	}

	data, err := ctx.Mod.Format()
	if err != nil {
		return nil, err
	}

	path := project.GetGoModFile(ctx.ProjectDir)
	oldData, _ := os.ReadFile(path)

	return []pipeline.Patch{
		patches.WriteFilePatch{
			Path:    path,
			OldData: oldData,
			NewData: data,
			Perm:    0666,
		},
	}, nil
}
