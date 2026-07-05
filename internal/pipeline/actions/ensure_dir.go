package actions

import (
	"os"

	"github.com/80LK/godev/internal/pipeline/context"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/utils"
)

type EnsureDir struct {
	Path string
	Perm os.FileMode
}

func (a EnsureDir) Plan(
	ctx *context.Context,
) ([]pipeline.Patch, error) {

	ok, err := utils.ExsistDir(a.Path)
	if err != nil {
		return nil, err
	}

	if ok {
		return nil, nil
	}

	return []pipeline.Patch{
		patches.CreateDirPatch{
			Path: a.Path,
			Perm: a.Perm,
		},
	}, nil
}
