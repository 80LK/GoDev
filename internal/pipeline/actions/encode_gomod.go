package actions

import (
	"fmt"
	"os"

	"github.com/80LK/godev/internal/pipeline/context"

	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/project"
)

type EncodeGoMod struct {
}

func (a EncodeGoMod) Plan(
	ctx *context.Context,
) ([]patches.Patch, error) {

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

	return []patches.Patch{
		patches.NewWriteFilePatch(
			path,
			oldData,
			data,
			0666,
		),
	}, nil
}
