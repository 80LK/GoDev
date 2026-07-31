package actions

import (
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type Action interface {
	Plan(*context.Context) ([]patches.Patch, error)
}

type Executor interface {
	Action
	Add(...Action) Executor
}
