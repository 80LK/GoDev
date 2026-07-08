package patches

import "github.com/80LK/godev/internal/pipeline/patches/context"

type Patch interface {
	Apply() error
	Summary(ctx *context.Context) (string, error)
}
