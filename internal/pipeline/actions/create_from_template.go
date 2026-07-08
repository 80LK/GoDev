package actions

import (
	"fmt"
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

type CreateFromTemplate struct {
	Template string
}

func (c CreateFromTemplate) Plan(ctx *context.Context) ([]patches.Patch, error) {
	switch c.Template {
	case "app":
		return createApp(ctx)
	case "module":
		return createModule(ctx)
	default:
		return nil, fmt.Errorf("Unknown template %s", c.Template)
	}
}

func createApp(ctx *context.Context) ([]patches.Patch, error) {
	return Pipeline{
		Actions: []Action{
			EnsureDir{Path: filepath.Join(ctx.ProjectDir, "cmd"), Perm: 0666},
			EnsureDir{Path: filepath.Join(ctx.ProjectDir, "cmd", ctx.GoProject.Project.Name), Perm: 0666},
			WriteFile{Path: filepath.Join(ctx.ProjectDir, "cmd", ctx.GoProject.Project.Name, ctx.GoProject.Project.Name+".go"), Perm: 0777, Value: []byte("package main\n\nfunc main(){\n}\n")},

			EnsureDir{Path: filepath.Join(ctx.ProjectDir, "internal"), Perm: 0666},
			EncodeGoProject{},
			EncodeGoMod{},
		},
	}.Plan(ctx)
}
func createModule(ctx *context.Context) ([]patches.Patch, error) {
	return Pipeline{
		Actions: []Action{
			WriteFile{Path: filepath.Join(ctx.ProjectDir, ctx.GoProject.Project.Name+".go"), Perm: 0777, Value: []byte(fmt.Sprintf("package %s\n", ctx.GoProject.Project.Name))},

			EnsureDir{Path: filepath.Join(ctx.ProjectDir, "internal"), Perm: 0666},
			EncodeGoProject{},
			EncodeGoMod{},
		},
	}.Plan(ctx)
}
