package pipeline

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline/actions"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
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
	i := 1
	for _, patch := range patches {
		sum := patch.Summary()
		diff, err := patch.Diff()

		if err != nil {
			return err
		}
		if sum == "" && diff == "" {
			continue
		}

		fmt.Printf("%d. %s\n", i, sum)
		i++
		if diff != "" {
			fmt.Println(diff)
		}
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
