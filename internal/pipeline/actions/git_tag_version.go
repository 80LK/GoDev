package actions

import (
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
	if err := cmd.Run(); err != nil {
		return nil, err
	}

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
