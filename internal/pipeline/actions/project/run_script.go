package actions

import (
	"fmt"
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type RunScript struct {
	Name           string
	IgnoreNotFound bool
}

func (r RunScript) Plan(ctx *context.Context) ([]patches.Patch, error) {
	script, ok := ctx.GoProject.Scripts[r.Name]
	if !ok {
		if r.IgnoreNotFound {
			return nil, nil
		} else {
			return nil, fmt.Errorf("Not found script %q", r.Name)
		}
	}

	workdir := ctx.ProjectDir
	if script.WorkDir != nil {
		workdir = filepath.Clean(filepath.Join(workdir, *script.WorkDir))
	}

	if len(script.Env) > 0 {
		fmt.Printf("Env now ignored")
	}

	return []patches.Patch{
		patches.ShellPatch{
			Command: script.Command,
			Args:    script.Args,
			WorkDir: workdir,

			Stdout: true,
		},
	}, nil
}
