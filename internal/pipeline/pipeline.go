package pipeline

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline/context"
)

type Pipeline struct {
	actions []Action
}

func New() *Pipeline {
	return &Pipeline{
		actions: make([]Action, 0, 4),
	}
}

func (p Pipeline) Plan(
	ctx *context.Context,
) ([]Patch, error) {

	var patches []Patch

	for _, action := range p.actions {

		actionPatches, err := action.Plan(ctx)
		if err != nil {
			return nil, err
		}

		patches = append(
			patches,
			actionPatches...,
		)
	}

	return patches, nil
}

func (p *Pipeline) Add(actions ...Action) *Pipeline {
	p.actions = append(p.actions, actions...)
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
	patches []Patch,
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
		if diff != "" {
			fmt.Println(diff)
		}
	}
	return nil
}

func apply(
	patches []Patch,
) error {

	for _, patch := range patches {
		if err := patch.Apply(); err != nil {
			return err
		}
	}

	return nil
}

type Action interface {
	Plan(*context.Context) ([]Patch, error)
}

type Patch interface {
	Apply() error
	Diff() (string, error)
	Summary() string
}
