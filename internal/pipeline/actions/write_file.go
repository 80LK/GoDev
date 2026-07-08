package actions

import (
	"fmt"
	"os"

	"github.com/80LK/godev/internal/pipeline/context"

	"github.com/80LK/godev/internal/pipeline/patches"
)

type WriteFile struct {
	Path string

	InputKey string
	Value    []byte

	Perm os.FileMode
}

func (a WriteFile) Plan(
	ctx *context.Context,
) ([]patches.Patch, error) {

	if a.Value == nil && a.InputKey == "" {
		return nil, fmt.Errorf("No such data")
	}

	var data []byte
	if a.InputKey != "" {
		var ok bool
		if data, ok = context.Get[[]byte](
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

	return []patches.Patch{
		patches.NewWriteFilePatch(
			a.Path,
			oldData,
			data,
			a.Perm,
		),
	}, nil
}
