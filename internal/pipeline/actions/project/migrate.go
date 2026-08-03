package actions

import (
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/project"
)

type Migrate struct{}

func (m Migrate) Plan(ctx *context.Context) ([]patches.Patch, error) {
	pchs := []patches.Patch{}
	switch ctx.GoProject.ConfigVersion {
	case 0:
		ctx.DeprecetedConfig = false
		ctx.GoProject.ConfigVersion = 1
		_p, err := EncodeGoProject{}.Plan(ctx)
		if err != nil {
			return nil, err
		}
		pchs = append(pchs, _p...)
		pchs = append(pchs, patches.MovePatch{
			Input:  project.GetGoProjectFile(ctx.ProjectDir),
			Output: filepath.Join(ctx.ProjectDir, "go.project.bck"),
		})
		fallthrough
	default:
		return pchs, nil
	}
}
