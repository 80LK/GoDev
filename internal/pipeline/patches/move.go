package patches

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline/patches/context"
)

type MovePatch struct {
	Input  string
	Output string
}

func (m MovePatch) Apply() error {
	if err := os.MkdirAll(filepath.Dir(m.Output), 0777); err != nil && !os.IsExist(err) {
		return err
	}
	return os.Rename(m.Input, m.Output)
}

func (m MovePatch) Summary(ctx *context.Context) (string, error) {
	ctx = context.Get(ctx)

	return fmt.Sprintf("%s%smove from %q to %q", ctx.GetPrefix(), ctx.GetCounter(), m.Input, m.Output), nil
}
