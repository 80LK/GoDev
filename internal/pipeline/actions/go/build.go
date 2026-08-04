package actions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/utils"
	"github.com/80LK/godev/internal/utils/logger"
	"github.com/80LK/godev/project"
)

type Build struct {
	Target string
}

func (b Build) Plan(ctx *context.Context) ([]patches.Patch, error) {
	logger := logger.Get("BUILD")
	if ctx.GoProject.Builds == nil {
		return nil, nil
	}

	build, ok := ctx.GoProject.Builds[b.Target]
	logger.Log("Build %q exist: %t", b.Target, ok)
	if !ok {
		fmt.Printf("Target %s not found. skiped\n", b.Target)
		return nil, nil
	}
	pchs := make([]patches.Patch, 0)

	logger.Log("Build %q pre-script: %q", b.Target, build.PreScript)
	if build.PreScript != "" {
		script, err := projectAct.RunScript{Name: build.PreScript}.Plan(ctx)
		if err != nil {
			return nil, err
		}
		pchs = append(pchs, script...)
	}

	pchs = append(pchs, buildShell(build, ctx.ProjectDir)...)

	logger.Log("Build %q post-script: %q", b.Target, build.PostScript)
	if build.PostScript != "" {
		script, err := projectAct.RunScript{Name: build.PostScript}.Plan(ctx)
		if err != nil {
			return nil, err
		}
		pchs = append(pchs, script...)
	}

	return pchs, nil
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

func buildShell(buildInfo *project.BuildInfo, wd string) []patches.Patch {
	logger := logger.Get("BUILD SHELL")
	patch := patches.ShellPatch{
		WorkDir: wd,
		Command: "go",
		Args: []string{
			"build",
		},
		Stdout: true,
		Env:    os.Environ(),
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

	tempOut := utils.GetTempFile(filepath.Base(buildInfo.Input))
	patch.Args = append(patch.Args,
		"-o",
		tempOut,
		buildInfo.Input,
	)
	logger.Log("tempOut")
	return []patches.Patch{
		patch,
		patches.MovePatch{
			Input:  tempOut,
			Output: buildInfo.Output,
		},
	}
}
