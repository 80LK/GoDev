package actions

import (
	"fmt"
	"strings"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/project"
)

type Builds struct {
	Target string
}

func appendBoolFlag(args []string, enabled bool, flag string) []string {
	if enabled {
		return append(args, flag)
	}

	return args
}

func appendListFlag(args []string, list []string, flag string, join string) []string {
	if len(list) > 0 {
		return append(args, flag+"="+strings.Join(list, join))
	}

	return args
}

func buildShell(buildInfo *project.BuildInfo, wd string) patches.ShellPatch {
	patch := patches.ShellPatch{
		WorkDir: wd,
		Command: "go",
		Args: []string{
			"build",
		},
	}

	patch.Args = appendBoolFlag(patch.Args, buildInfo.Race, "-race")
	patch.Args = appendBoolFlag(patch.Args, buildInfo.Trimpath, "-trimpath")

	patch.Args = appendListFlag(patch.Args, buildInfo.Tags, "-tags", ",")
	patch.Args = appendListFlag(patch.Args, buildInfo.LdFlags, "-ldflags", " ")
	patch.Args = appendListFlag(patch.Args, buildInfo.GcFlags, "-gcflags", " ")

	patch.Args = append(patch.Args,
		"-o",
		buildInfo.Output,
		buildInfo.Input,
	)

	return patch
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
			ptchs.Items[i] = buildShell(build, ctx.ProjectDir)
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
		buildShell(build, ctx.ProjectDir),
	}, nil
}
