package actions

import (
	"fmt"
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline"
)

type CreateFromTemplate struct {
	Template string
}

func (c CreateFromTemplate) Plan(ctx *pipeline.Context) ([]pipeline.Patch, error) {
	switch c.Template {
	case "app":
		return createApp(ctx)
	case "module":
		return createModule(ctx)
	default:
		return nil, fmt.Errorf("Unknown template %s", c.Template)
	}
}

func createApp(ctx *pipeline.Context) ([]pipeline.Patch, error) {
	return pipeline.New().Add(
		EnsureDir{Path: filepath.Join(ctx.ProjectDir, "cmd"), Perm: 0666},
		EnsureDir{Path: filepath.Join(ctx.ProjectDir, "cmd", ctx.Project.Project.Name), Perm: 0666},
		WriteFile{Path: filepath.Join(ctx.ProjectDir, "cmd", ctx.Project.Project.Name, ctx.Project.Project.Name+".go"), Perm: 0777, Value: []byte("package main\n\nfunc main(){\n}\n")},

		EnsureDir{Path: filepath.Join(ctx.ProjectDir, "internal"), Perm: 0666},
		EncodeGoProject{},
		EncodeGoMod{},
	).Plan(ctx)
}
func createModule(ctx *pipeline.Context) ([]pipeline.Patch, error) {
	return pipeline.New().Add(
		WriteFile{Path: filepath.Join(ctx.ProjectDir, ctx.Project.Project.Name+".go"), Perm: 0777, Value: []byte(fmt.Sprintf("package %s\n", ctx.Project.Project.Name))},

		EnsureDir{Path: filepath.Join(ctx.ProjectDir, "internal"), Perm: 0666},
		EncodeGoProject{},
		EncodeGoMod{},
	).Plan(ctx)
}
