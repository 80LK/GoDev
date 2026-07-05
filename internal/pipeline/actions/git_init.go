package actions

import (
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/utils"
)

type GitInit struct{}

func (f GitInit) Plan(ctx *pipeline.Context) ([]pipeline.Patch, error) {
	exsist, err := utils.ExsistDir(filepath.Join(ctx.ProjectDir, ".git"))
	if err != nil {
		return nil, err
	}

	if exsist {
		return nil, nil
	}

	return []pipeline.Patch{
		patches.ShellPatch{
			Command: "git",
			Args:    []string{"init", "."},
			WorkDir: ctx.ProjectDir,
		},
	}, nil
}
