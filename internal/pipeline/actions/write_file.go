package actions

import (
	"fmt"
	"os"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type WriteFile struct {
	Path string

	InputKey string
	Value    []byte

	Perm os.FileMode
}

func (a WriteFile) Plan(
	ctx *pipeline.Context,
) ([]pipeline.Patch, error) {

	if a.Value == nil && a.InputKey == "" {
		return nil, fmt.Errorf("No such data")
	}

	var data []byte
	if a.InputKey != "" {
		var ok bool
		if data, ok = pipeline.Get[[]byte](
			ctx,
			a.InputKey,
		); !ok {
			return nil,
				fmt.Errorf(
					"context key %q not found",
					a.InputKey,
				)
		}
	} else {
		data = a.Value
	}

	oldData, _ := os.ReadFile(a.Path)

	return []pipeline.Patch{
		patches.WriteFilePatch{
			Path:    a.Path,
			OldData: oldData,
			NewData: data,
			Perm:    a.Perm,
		},
	}, nil
}
