package patches

import (
	"os"

	"github.com/80LK/godev/internal/pipeline/patches/context"
)

type CreateDirPatch struct {
	Path string
	Perm os.FileMode
}

func (p CreateDirPatch) Apply() error {
	return os.MkdirAll(p.Path, p.Perm)
}

func (p CreateDirPatch) Summary(ctx *context.Context) (string, error) {
	ctx = context.Get(ctx)

	return ctx.GetPrefix() + ctx.GetCounter() + "mkdir " + p.Path + "\n", nil
}
