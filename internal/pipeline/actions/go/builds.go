package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type Builds struct{}

func (b Builds) Plan(ctx *context.Context) ([]patches.Patch, error) {
	if ctx.GoProject.Builds == nil {
		return nil, nil
	}

	ptchs := patches.ParallelPatch{
		Items: make([]patches.Patch, len(ctx.GoProject.Builds)),
	}

	for i, build := range ctx.GoProject.Builds {
		ptchs.Items[i] = patches.ShellPatch{
			Command: "go",
			Args: []string{
				"build",
				"-o",
				build.Output,
				build.Input,
			},
			WorkDir: ctx.ProjectDir,
		}
	}

	return []patches.Patch{
		ptchs,
	}, nil
}
