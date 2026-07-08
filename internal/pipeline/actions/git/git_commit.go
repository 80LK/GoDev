package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type GitCommit struct {
	InputKey string
	Value    string
}

func (g GitCommit) Plan(*context.Context) ([]patches.Patch, error) {
	panic("unimplemented")
}
