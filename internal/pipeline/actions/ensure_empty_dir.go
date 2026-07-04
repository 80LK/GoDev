package actions

import (
	"fmt"
	"os"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/utils"
)

type EnsureEmptyDir struct {
	Path string
	Perm os.FileMode
}

func (a EnsureEmptyDir) Plan(
	ctx *pipeline.Context,
) ([]pipeline.Patch, error) {

	ok, err := utils.ExsistDir(a.Path)
	if err != nil {
		return nil, err
	}

	if !ok {
		return []pipeline.Patch{
			patches.CreateDirPatch{
				Path: a.Path,
				Perm: a.Perm,
			},
		}, nil
	}

	dirs, err := os.ReadDir(a.Path)
	if len(dirs) > 0 {
		return nil, fmt.Errorf("%s not empty", a.Path)
	}

	return nil, nil
}
