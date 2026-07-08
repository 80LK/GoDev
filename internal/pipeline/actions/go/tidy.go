package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type Tidy struct{}

func (t Tidy) Plan(ctx *context.Context) ([]patches.Patch, error) {
	return []patches.Patch{
		patches.ShellPatch{
			Command: "go",
			Args:    []string{"mod", "tidy"},
			WorkDir: ctx.ProjectDir,
		},
	}, nil
}
