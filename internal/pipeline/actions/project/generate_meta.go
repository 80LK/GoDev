package actions

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"text/template"

	_ "embed"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
)

//go:embed templates/generated.go.tpl
var templateMetaSource string

type _GenerateMeta struct {
	tpl template.Template
}

func NewGenerateMeta() (*_GenerateMeta, error) {
	tpl, err := template.New("generated.go").Parse(templateMetaSource)
	if err != nil {
		return nil, err
	}

	return &_GenerateMeta{
		tpl: *tpl,
	}, nil
}

func (g _GenerateMeta) Plan(ctx *context.Context) ([]patches.Patch, error) {

	var buf bytes.Buffer

	if err := g.tpl.Execute(&buf, ctx.GoProject.Project); err != nil {
		return nil, err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}

	path := filepath.Join(ctx.ProjectDir, "internal", "meta", "generated.go")

	oldData, _ := os.ReadFile(path)

	return []patches.Patch{
		patches.NewWriteFilePatch(path, oldData, formatted, 0666),
	}, nil
}
