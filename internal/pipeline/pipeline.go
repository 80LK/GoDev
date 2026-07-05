package pipeline

import (
	"fmt"
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
	ctx *Context,
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

func (p *Pipeline) Execute(ctx *Context) error {
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
	for i, patch := range patches {
		sum := patch.Summary()
		fmt.Printf("%d. %s\n", i+1, sum)

		diff, err := patch.Diff()
		if err != nil {
			return err
		}
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
	Plan(*Context) ([]Patch, error)
}

type Patch interface {
	Apply() error
	Diff() (string, error)
	Summary() string
}
