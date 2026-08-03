package actions

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/modlike"

	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/project"
)

type EncodeGoProject struct {
}

func (a EncodeGoProject) Plan(
	ctx *context.Context,
) ([]patches.Patch, error) {

	if ctx.GoProject == nil {
		return nil,
			fmt.Errorf(
				"project not found in context",
			)
	}

	var (
		data []byte
		err  error
		path string
	)
	if ctx.DeprecetedConfig {
		data, err = modlike.Marshal(ctx.GoProject)
		if err != nil {
			return nil, err
		}
		path = project.GetGoProjectFile(ctx.ProjectDir)
	} else {
		data, err = json.Marshal(ctx.GoProject)
		if err != nil {
			return nil, err
		}
		path = project.GetJSONProjectFile(ctx.ProjectDir)
	}

	oldData, _ := os.ReadFile(path)

	return []patches.Patch{
		patches.NewWriteFilePatch(
			path,
			oldData,
			data,
			0666,
		),
	}, nil
}
