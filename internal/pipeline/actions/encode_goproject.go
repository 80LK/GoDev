package actions

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline"

	"github.com/80lk/modlike"
)

type EncodeGoProject struct {
	OutputKey string
}

func (a EncodeGoProject) Plan(
	ctx *pipeline.Context,
) ([]pipeline.Patch, error) {

	if ctx.Project == nil {
		return nil,
			fmt.Errorf(
				"project not found in context",
			)
	}

	doc := modlike.New()

	err := doc.Encode(ctx.Project)
	if err != nil {
		return nil, err
	}

	data, err :=
		modlike.Serialize(doc)

	if err != nil {
		return nil, err
	}

	pipeline.Set(
		ctx,
		a.OutputKey,
		data,
	)

	return nil, nil
}
