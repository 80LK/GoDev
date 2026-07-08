package actions

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"

	"github.com/80LK/godev/internal/utils"
)

type CheckExistsFile struct {
	Path  string
	Error string
}

func (a CheckExistsFile) Plan(*context.Context) ([]patches.Patch, error) {
	if a.Error == "" {
		a.Error = "%s not exists"
	}

	ok, err := utils.ExsistFile(a.Path)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf(a.Error, a.Path)
	}

	return nil, nil
}
