package actions

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type Builds struct{}

func (b Builds) Plan(ctx *context.Context) ([]pipeline.Patch, error) {
	if ctx.GoProject.Builds == nil {
		return nil, nil
	}

	ptchs := patches.ParallelPatch{
		Items: make([]pipeline.Patch, len(ctx.GoProject.Builds)),
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
		}
	}

	return []pipeline.Patch{
		ptchs,
	}, nil
}
