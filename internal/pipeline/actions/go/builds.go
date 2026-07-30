package actions

import (
	"fmt"
	"os"
	"strings"

	"github.com/80LK/godev/internal/pipeline/actions"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/project"
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

	if b.Target == "" {
		paralelPlan := actions.Parallel{
			Items: make([]actions.Action, len(ctx.GoProject.Builds)),
		}

		i := 0
		for name := range ctx.GoProject.Builds {
			paralelPlan.Items[i] = Builds{Target: name}
			i++
		}

		return paralelPlan.Plan(ctx)

	}

	build, ok := ctx.GoProject.Builds[b.Target]
	if !ok {
		return nil, fmt.Errorf("Target %s not found", b.Target)
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

	return pchs, nil
}
