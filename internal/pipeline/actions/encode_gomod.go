package actions

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline"

	"golang.org/x/mod/modfile"
)

type EncodeGoMod struct {
	OutputKey string
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

	pipeline.Set(
		ctx,
		a.OutputKey,
		data,
	)

	return nil, nil
}
