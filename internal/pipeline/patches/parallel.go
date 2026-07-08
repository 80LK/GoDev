package patches

import (
	"strings"

	"github.com/80LK/godev/internal/pipeline/patches/context"
	"golang.org/x/sync/errgroup"
)

type ParallelPatch struct {
	Items []Patch
}

func (p ParallelPatch) Apply() error {
	l := len(p.Items)
	switch l {
	case 0:
		return nil
	case 1:
		return p.Items[0].Apply()
	}

	var g errgroup.Group

	for _, item := range p.Items {
		g.Go(func() error {
			return item.Apply()
		})
	}

	return g.Wait()
}
func (p ParallelPatch) Summary(ctx *context.Context) (string, error) {
	ctx = context.Get(ctx)

	l := len(p.Items)
	switch l {
	case 0:
		return "", nil
	case 1:
		return p.Items[0].Summary(ctx)
	}

	var str strings.Builder

	str.WriteString(ctx.GetPrefix() + ctx.GetCounter() + "Parallel:\n")
	nextLevelCtx := ctx.NextLevel()
	for _, item := range p.Items {
		sum, err := item.Summary(nextLevelCtx)
		if err != nil {
			return "", err
		}
		if sum == "" {
			continue
		}

		str.WriteString(sum)
	}

	return str.String(), nil
}
