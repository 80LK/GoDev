package actions

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type GitTagVersion struct{}

func (g GitTagVersion) Plan(ctx *context.Context) ([]pipeline.Patch, error) {
	tag := ctx.GoProject.Project.Version.String()

	cmd := exec.Command("git", "rev-parse", "-q", "--verify", "refs/tags/"+tag)
	cmd.Dir = ctx.ProjectDir
	err := cmd.Run()
	if err == nil {
		return nil, fmt.Errorf("tag already exsist. change version")
	}

	_e, ok := errors.AsType[*exec.ExitError](err)
	if ok {
		if _e.ExitCode() == 1 {
			return []pipeline.Patch{
				patches.ShellPatch{
					Command: "git",
					Args: []string{
						"tag",
						"-a",
						tag,
						"-m",
						"Release " + tag,
					},
					WorkDir: ctx.ProjectDir,
				},
			}, nil
		}
	}

	return nil, err
}
