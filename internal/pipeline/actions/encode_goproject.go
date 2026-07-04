package actions

import (
	"fmt"
	"os"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/project"

	"github.com/80lk/modlike"
)

type EncodeGoProject struct {
}

func (a EncodeGoProject) Plan(
	ctx *pipeline.Context,
) ([]pipeline.Patch, error) {

	if ctx.Project == nil {
		return nil,
			fmt.Errorf(
				"project not found in context",
			)
	}

	doc := modlike.New()

	err := doc.Encode(ctx.Project)
	if err != nil {
		return nil, err
	}

	data, err :=
		modlike.Serialize(doc)

	if err != nil {
		return nil, err
	}

	path := project.GetGoProjectFile(ctx.ProjectDir)
	oldData, _ := os.ReadFile(path)

	return []pipeline.Patch{
		patches.WriteFilePatch{
			Path:    path,
			OldData: oldData,
			NewData: data,
			Perm:    0666,
		},
	}, nil
}
