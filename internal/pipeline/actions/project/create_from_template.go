package actions

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/project"
)

//go:embed all:templates/app
var templateApp embed.FS

//go:embed all:templates/module
var templateModule embed.FS

type CreateFromTemplate struct {
	Template string
}

func (c CreateFromTemplate) Plan(ctx *context.Context) ([]patches.Patch, error) {
	switch c.Template {
	case "app":
		return generateTemplate(ctx, &templateApp, "templates/app")
	case "module":
		return generateTemplate(ctx, &templateModule, "templates/module")
	default:
		return nil, fmt.Errorf("Unknown template %s", c.Template)
	}
}

type TemplateContext struct {
	Project   *project.ProjectInfo
	GoVersion string
}

func generateTemplate(ctx *context.Context, templateFS *embed.FS, root string) ([]patches.Patch, error) {
	tplContext := &TemplateContext{
		Project:   ctx.GoProject.Project,
		GoVersion: ctx.GoVer.StringWithoutSuffix(),
	}

	ptchs := []patches.Patch{}
	err := fs.WalkDir(templateFS, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		tpl, err := template.New("Path name: " + rel).Parse(rel)
		if err != nil {
			return err
		}

		var relPathBuilder strings.Builder
		if err := tpl.Execute(&relPathBuilder, tplContext); err != nil {
			return err
		}
		rel = relPathBuilder.String()

		targetPath := filepath.Join(ctx.ProjectDir, rel)

		if d.IsDir() {
			ptchs = append(ptchs, patches.CreateDirPatch{
				Path: targetPath,
				Perm: 0666,
			})
			return nil
		}

		data, err := fs.ReadFile(templateFS, path)
		if err != nil {
			return err
		}

		if trimmedPath, ok := strings.CutSuffix(targetPath, ".tpl"); ok {
			targetPath = trimmedPath
			tpl, err := template.New("File: " + targetPath).Parse(string(data))
			if err != nil {
				return err
			}

			var buf bytes.Buffer
			if err := tpl.Execute(&buf, tplContext); err != nil {
				return err
			}
			data = buf.Bytes()
		}

		ptchs = append(ptchs, patches.NewWriteFilePatch(targetPath, nil, data, 0777))

		return nil
	})

	return ptchs, err
}
