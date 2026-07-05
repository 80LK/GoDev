package actions

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/context"
)

type CheckClearGit struct{}

func (c CheckClearGit) Plan(ctx *context.Context) ([]pipeline.Patch, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = ctx.ProjectDir

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	if len(bytes.TrimSpace(out)) != 0 {
		return nil, fmt.Errorf("git repository is not clean")
	}

	return nil, nil
}
