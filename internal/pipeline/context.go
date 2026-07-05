package pipeline

import (
	"os/exec"

	"github.com/80LK/godev/internal/project"
	"github.com/80LK/godev/internal/version"
	"golang.org/x/mod/modfile"
)

type Context struct {
	DryRun     bool
	ProjectDir string

	Project *project.GoProject
	Mod     *modfile.File

	GoVer *version.Version

	values map[string]any
}

func NewContext(dryRun bool) *Context {
	out, err := exec.Command("go", "env", "GOVERSION").Output()
	if err != nil {
		panic(err)
	}
	ver, err := version.Parse(string(out))
	if err != nil {
		panic(err)
	}

	return &Context{
		DryRun: dryRun,
		GoVer:  ver,
		values: make(map[string]any),
	}
}

func Set[T any](ctx *Context, key string, value T) {
	ctx.values[key] = value
}

func Get[T any](ctx *Context, key string) (T, bool) {
	value, ok := ctx.values[key]
	if !ok {
		var zero T
		return zero, false
	}

	result, ok := value.(T)
	if !ok {
		var zero T
		return zero, false
	}

	return result, true
}
