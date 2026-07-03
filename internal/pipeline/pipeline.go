package pipeline

import (
	"fmt"
)

type Pipeline struct {
	Actions []Action
}

type Action interface {
	Plan(*Context) ([]Patch, error)
}

type Patch interface {
	Apply() error
	Diff() (string, error)
	Summary() string
}

func (p Pipeline) Plan(
	ctx *Context,
) ([]Patch, error) {

	var patches []Patch

	for _, action := range p.Actions {

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

func Execute(ctx *Context, p *Pipeline) error {
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

	for _, patch := range patches {
		diff, err := patch.Diff()
		if err != nil {
			return err
		}

		fmt.Println(diff)
	}

	fmt.Println("Summary:")
	i := 1
	for _, patch := range patches {
		sum := patch.Summary()
		if sum != "" {
			fmt.Printf("%d. %s\n", i, sum)
			i++
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
