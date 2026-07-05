package actions

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/version"
)

type PatchSources struct{}

const _OLD_VERSION_CONTEXT_KEY = "old_version"

func (p PatchSources) Plan(ctx *context.Context) ([]pipeline.Patch, error) {
	oldVersion, ok := context.Get[*version.Version](ctx, _OLD_VERSION_CONTEXT_KEY)
	if !ok {
		return nil, nil
	}

	if ctx.GoProject.Project.Version.Major < 2 {
		return nil, nil
	}

	if oldVersion.Major == ctx.GoProject.Project.Version.Major {
		return nil, nil
	}

	oldModule := ctx.GoProject.Project.Module
	newModule := ctx.GoProject.Project.Module + "/v" + strconv.FormatUint(uint64(ctx.GoProject.Project.Version.Major), 10)

	// patch go.mod
	ctx.Mod.AddModuleStmt(newModule)
	encode := EncodeGoMod{}
	patchs, err := encode.Plan(ctx)
	if err != nil {
		return nil, err
	}

	// patch *.go

	err = filepath.WalkDir(ctx.ProjectDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == ".go" {
			oldData, _ := os.ReadFile(path)

			fset := token.NewFileSet()

			file, err := parser.ParseFile(
				fset,
				path,
				oldData,
				parser.ParseComments,
			)
			if err != nil {
				return err
			}

			for _, imp := range file.Imports {
				path, _ := strconv.Unquote(imp.Path.Value)
				if path == oldModule {
					imp.Path.Value = strconv.Quote(newModule)
				} else if v, ok := strings.CutPrefix(path, oldModule+"/"); ok {
					path = newModule + "/" + v
				} else {
					continue
				}
				imp.Path.Value = strconv.Quote(path)
			}

			var buf bytes.Buffer
			if err := format.Node(&buf, fset, file); err != nil {
				return err
			}

			newData := buf.Bytes()

			patchs = append(patchs, patches.WriteFilePatch{
				Path:    path,
				OldData: oldData,
				NewData: newData,
				Perm:    0644,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return patchs, nil
}
