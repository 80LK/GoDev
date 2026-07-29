package patches

import (
	"strings"

	"github.com/80LK/godev/internal/pipeline/patches/context"
	"golang.org/x/sync/errgroup"
)

type ParallelPatch struct {
	Items [][]Patch
}

func (p ParallelPatch) Apply() error {
	l := len(p.Items)
	switch l {
	case 0:
		return nil
	case 1:
		for _, patch := range p.Items[0] {
			err := patch.Apply()
			if err != nil {
				return err
			}
		}
		return nil
	}

	var g errgroup.Group

	for _, item := range p.Items {
		g.Go(func() error {
			for _, patch := range item {
				err := patch.Apply()
				if err != nil {
					return err
				}
			}
			return nil
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
		var str strings.Builder
		for _, patch := range p.Items[0] {
			sum, err := patch.Summary(ctx)
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

	var str strings.Builder

	str.WriteString(ctx.GetPrefix())
	str.WriteString(ctx.GetCounter())
	str.WriteString("Parallel:\n")
	nextLevelCtx := ctx.NextLevel()
	for _, item := range p.Items {
		for _, patch := range item {
			sum, err := patch.Summary(nextLevelCtx)
			if err != nil {
				return "", err
			}
			if sum == "" {
				continue
			}

			str.WriteString(sum)
		}

	}

	return str.String(), nil
}
