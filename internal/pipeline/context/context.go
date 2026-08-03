package context

import (
	"os"
	"os/exec"

	"github.com/80LK/godev/internal/version"
	"github.com/80LK/godev/project"
	"golang.org/x/mod/modfile"
)

type Context struct {
	DryRun     bool
	ProjectDir string

	GoProject        *project.GoProject
	DeprecetedConfig bool
	Mod              *modfile.File

	GoVer *version.Version

	values map[string]any
}

func New(dryRun bool) *Context {
	out, err := exec.Command("go", "env", "GOVERSION").Output()
	if err != nil {
		panic(err)
	}
	ver, err := version.Parse(string(out))
	if err != nil {
		panic(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return &Context{
		DryRun:     dryRun,
		ProjectDir: cwd,
		GoVer:      ver,
		values:     make(map[string]any),
		GoProject:  nil,
		Mod:        nil,
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
