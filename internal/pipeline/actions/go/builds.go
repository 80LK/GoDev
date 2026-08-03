package actions

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/80LK/godev/internal/pipeline/actions"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/project"
)

type Builds struct {
	Targets  []string
	Parallel bool
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
		Env: os.Environ(),
	}

	patch.Args = appendBoolFlag(patch.Args, buildInfo.Race, "-race")
	patch.Args = appendBoolFlag(patch.Args, buildInfo.Trimpath, "-trimpath")

	patch.Args = appendListFlag(patch.Args, append(buildInfo.Tags, "god"), "-tags", ",")
	patch.Args = appendListFlag(patch.Args, buildInfo.LdFlags, "-ldflags", " ")
	patch.Args = appendListFlag(patch.Args, buildInfo.GcFlags, "-gcflags", " ")

	if buildInfo.OS != "" {
		patch.Env = append(patch.Env, "GOOS="+buildInfo.OS)
	}
	if buildInfo.Arch != "" {
		patch.Env = append(patch.Env, "GOARCH="+buildInfo.Arch)
	}

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

	var plan actions.Executor

	if b.Parallel {
		plan = &actions.Parallel{}
	} else {
		plan = &actions.Pipeline{}
	}

	if b.Targets == nil {
		b.Targets = slices.Collect(maps.Keys(ctx.GoProject.Builds))
	}

	for _, buildName := range b.Targets {
		build, ok := ctx.GoProject.Builds[buildName]
		if !ok {
			fmt.Printf("Target %s not found. skiped\n", b.Targets)
			continue
		}
		pchs := make([]patches.Patch, 0)

		if build.PreScript != "" {
			script, err := projectAct.RunScript{Name: build.PreScript}.Plan(ctx)
			if err != nil {
				return nil, err
			}
			pchs = append(pchs, script...)
		}
		pchs = append(pchs, buildShell(build, ctx.ProjectDir))
		if build.PostScript != "" {
			script, err := projectAct.RunScript{Name: build.PostScript}.Plan(ctx)
			if err != nil {
				return nil, err
			}
			pchs = append(pchs, script...)
		}

	}

	return plan.Plan(ctx)
}
