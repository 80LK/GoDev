package actions

import (
	"fmt"
	"os"

	"github.com/80LK/godev/internal/pipeline"
)

type CheckNotExists struct {
	Path  string
	Error string
}

func (a CheckNotExists) Plan(
	ctx *pipeline.Context,
) ([]pipeline.Patch, error) {
	if a.Error == "" {
		a.Error = "%s already exists"
	}

	_, err := os.Stat(a.Path)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	return nil, fmt.Errorf(a.Error, a.Path)
}
