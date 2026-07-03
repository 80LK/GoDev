package actions

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/utils"
)

type CheckNotExists struct {
	Path string
}

func (a CheckNotExists) Plan(
	ctx *pipeline.Context,
) ([]pipeline.Patch, error) {

	ok, err := utils.ExsistFile(a.Path)
	if err != nil {
		return nil, err
	}

	if ok {
		return nil,
			fmt.Errorf(
				"%s already exists",
				a.Path,
			)
	}

	return nil, nil
}
