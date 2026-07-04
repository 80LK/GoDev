package actions

import (
	"fmt"
	"os"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/project"

	"golang.org/x/mod/modfile"
)

type EncodeGoMod struct {
}

func (a EncodeGoMod) Plan(
	ctx *pipeline.Context,
) ([]pipeline.Patch, error) {

	if ctx.Project == nil {
		return nil,
			fmt.Errorf(
				"project not found in context",
			)
	}

	mod := &modfile.File{}

	err := mod.AddModuleStmt(ctx.Project.Project.Module)
	if err != nil {
		return nil, err
	}

	err = mod.AddGoStmt(ctx.GoVer.StringWithoutSuffix())
	if err != nil {
		return nil, err
	}

	data, err := mod.Format()
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
