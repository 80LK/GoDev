package actions

import (
	"bytes"
	"embed"
	"path"
	"path/filepath"
	"text/template"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/utils"
)

//go:embed templates/meta
var templateMetaSource embed.FS

const _ROOT_PATH = "templates/meta"

var _FILES = []struct {
	File    string
	Rewrite bool
}{
	{File: "meta.go"},
	{File: "generated.go", Rewrite: true},
}

type GenerateMeta struct{}

func (g GenerateMeta) Plan(ctx *context.Context) ([]patches.Patch, error) {
	ptchs := []patches.Patch{}

	metaDir := filepath.Join(ctx.ProjectDir, "internal/meta")
	ok, err := utils.ExsistDir(metaDir)
	if err != nil {
		return nil, err
	}

	if !ok {
		ptchs = append(ptchs, patches.CreateDirPatch{
			Path: metaDir,
			Perm: 0666,
		})
	}

	for _, file := range _FILES {
		target := filepath.Join(metaDir, file.File)

		if !file.Rewrite {
			ok, err := utils.ExsistFile(target)
			if err != nil {
				return nil, err
			}

			if ok {
				continue
			}
		}
		source := path.Join(_ROOT_PATH, file.File)

		data, err := templateMetaSource.ReadFile(source)
		if err != nil {
			return nil, err
		}

		tpl, err := template.New(source).Parse(string(data))
		if err != nil {
			return nil, err
		}

		var buf bytes.Buffer
		if err := tpl.Execute(&buf, ctx.GoProject.Project); err != nil {
			return nil, err
		}
		data = buf.Bytes()

		ptchs = append(ptchs, patches.NewWriteFilePatch(target, nil, data, 0777))
	}

	return ptchs, nil
}
