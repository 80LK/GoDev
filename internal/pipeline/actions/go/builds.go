package actions

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type Builds struct {
	Target string
}

func (b Builds) Plan(ctx *context.Context) ([]patches.Patch, error) {
	if ctx.GoProject.Builds == nil {
		return nil, nil
	}

	if b.Target == "" {
		ptchs := patches.ParallelPatch{
			Items: make([]patches.Patch, len(ctx.GoProject.Builds)),
		}

		i := 0
		for _, build := range ctx.GoProject.Builds {
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
			i++
		}

		return []patches.Patch{
			ptchs,
		}, nil
	}

	build, ok := ctx.GoProject.Builds[b.Target]
	if !ok {
		return nil, fmt.Errorf("Target %s not found", b.Target)
	}

	return []patches.Patch{
		patches.ShellPatch{
			Command: "go",
			Args: []string{
				"build",
				"-o",
				build.Output,
				build.Input,
			},
			WorkDir: ctx.ProjectDir,
		},
	}, nil
}
