package actions

import (
	"bytes"
	"embed"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/utils"
	"github.com/pmezard/go-difflib/difflib"
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

type GenerateMeta struct {
	ContextKey string
}

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
		exsist, err := utils.ExsistFile(target)
		if err != nil {
			return nil, err
		}
		if !file.Rewrite && exsist {
			continue
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

		oldData, err := os.ReadFile(target)
		if err != nil {
			return nil, err
		}

		aLines := difflib.SplitLines(string(oldData))
		bLines := difflib.SplitLines(string(data))

		m := difflib.NewMatcher(aLines, bLines)

		if g.ContextKey != "" {
			context.Set(ctx, g.ContextKey, m.Ratio() != 1)
		}

		if m.Ratio() != 1 {
			ptchs = append(ptchs, patches.NewWriteFilePatch(target, oldData, data, 0777))
		}
	}

	return ptchs, nil
}
