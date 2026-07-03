package actions

import (
	"os/user"
	"path"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/version"
)

type InitProject struct {
	ModuleName string
}

func (f InitProject) Plan(ctx *pipeline.Context) ([]pipeline.Patch, error) {
	p := *ctx.Project
	if p.Project.Name == "" {
		p.Project.Name = path.Base(f.ModuleName)
	}

	if p.Project.Author == "" {
		user, err := user.Current()
		if err != nil {
			return nil, err
		}

		p.Project.Author = user.Name
	}

	if p.Project.Module == "" {
		p.Project.Module = f.ModuleName
	}

	if p.Project.Version == nil {
		p.Project.Version, _ = version.Parse("0.1.0")
	}

	return nil, nil
}
