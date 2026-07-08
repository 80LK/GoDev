package pipeline

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline/actions"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	summary "github.com/80LK/godev/internal/pipeline/patches/context"
)

type Pipeline struct {
	actions.Pipeline
}

func New() *Pipeline {
	p := &Pipeline{}
	p.Actions = make([]actions.Action, 0, 4)
	return p
}

func (p *Pipeline) Add(actions ...actions.Action) *Pipeline {
	p.Pipeline.Add(actions...)
	return p
}

func (p *Pipeline) Execute(ctx *context.Context) error {
	patches, err := p.Plan(ctx)
	if err != nil {
		return err
	}

	if ctx.DryRun {
		return dryRun(patches)
	}

	return apply(patches)
}

func dryRun(
	patches []patches.Patch,
) error {
	ctx := summary.Default()
	for _, patch := range patches {
		sum, err := patch.Summary(ctx)

		if err != nil {
			return err
		}
		if sum == "" {
			continue
		}

		fmt.Printf("%s", sum)
	}
	return nil
}

func apply(
	patches []patches.Patch,
) error {

	for _, patch := range patches {
		if err := patch.Apply(); err != nil {
			return err
		}
	}

	return nil
}
