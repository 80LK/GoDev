package actions

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type GitCommit struct {
	InputKey string
	Value    string
}

func (g GitCommit) Plan(ctx *context.Context) ([]patches.Patch, error) {
	if g.InputKey != "" {
		if _v, ok := context.Get[string](ctx, g.InputKey); ok {
			g.Value = _v
		} else if _v, ok := context.Get[fmt.Stringer](ctx, g.InputKey); ok {
			g.Value = _v.String()
		}
	}

	if g.Value == "" {
		return nil, fmt.Errorf("Can't get value for name commit")
	}

	return []patches.Patch{
		patches.ShellPatch{
			Command: "git",
			Args:    []string{"add", "."},
		},
		patches.ShellPatch{
			Command: "git",
			Args:    []string{"commit", "-m", g.Value},
		},
	}, nil
}
