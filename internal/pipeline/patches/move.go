package patches

import (
	"fmt"
	"os"

	"github.com/80LK/godev/internal/pipeline/patches/context"
)

type MovePatch struct {
	Input  string
	Output string
}

func (m MovePatch) Apply() error {
	return os.Rename(m.Input, m.Output)
}

func (m MovePatch) Summary(ctx *context.Context) (string, error) {
	ctx = context.Get(ctx)

	return fmt.Sprintf("%s%smove from %q to %q", ctx.GetPrefix(), ctx.GetCounter(), m.Input, m.Output), nil
}
