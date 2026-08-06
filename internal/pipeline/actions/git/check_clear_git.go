package actions

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type CheckClearGit struct {
	OutputKey string
}

func (c CheckClearGit) Plan(ctx *context.Context) ([]patches.Patch, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = ctx.ProjectDir

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	empty := len(bytes.TrimSpace(out)) != 0
	if c.OutputKey != "" {
		context.Set(ctx, c.OutputKey, empty)
	} else if !empty {
		return nil, fmt.Errorf("git repository is not clean")
	}

	return nil, nil
}
