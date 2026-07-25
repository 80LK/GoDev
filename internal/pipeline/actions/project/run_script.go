package actions

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/project"
)

type RunScript struct {
	Name           string
	IgnoreNotFound bool
}

func parseScriptInShellPatch(script *project.Script, workdir string, defaultScript *project.Script) (patches.Patch, error) {
	if script.Command == nil {
		return nil, fmt.Errorf("Command cant was been empty")
	}

	if script.WorkDir != nil {
		workdir = filepath.Clean(filepath.Join(workdir, *script.WorkDir))
	} else if defaultScript != nil && defaultScript.WorkDir != nil {
		workdir = filepath.Clean(filepath.Join(workdir, *defaultScript.WorkDir))
	}

	var env []string = script.Env
	if len(env) == 0 && defaultScript != nil {
		env = defaultScript.Env
	}

	return &patches.ShellPatch{
		Command: *script.Command,
		Args:    script.Args,
		WorkDir: workdir,
		Env:     env,

		Stdout: true,
	}, nil
}

func (r RunScript) Plan(ctx *context.Context) ([]patches.Patch, error) {
	script, ok := ctx.GoProject.Scripts[r.Name]
	if !ok {
		if r.IgnoreNotFound {
			return nil, nil
		} else {
			return nil, fmt.Errorf("Not found script %q", r.Name)
		}
	}

	if len(script.Commands) == 0 {
		log.Printf("Script: %+v\n", script)
		patch, err := parseScriptInShellPatch(script.AsScript(), ctx.ProjectDir, nil)
		if err != nil {
			return nil, err
		}
		return []patches.Patch{patch}, nil
	}

	ptchs := []patches.Patch{}

	for _, childScript := range script.Commands {
		patch, err := parseScriptInShellPatch(childScript, ctx.ProjectDir, script.AsScript())
		if err != nil {
			return nil, err
		}
		ptchs = append(ptchs, patch)
	}

	return ptchs, nil
}
