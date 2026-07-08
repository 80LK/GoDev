package patches

import (
	"os"
	"os/exec"
	"strings"

	"github.com/80LK/godev/internal/pipeline/patches/context"
)

type ShellPatch struct {
	Command string
	Args    []string
	WorkDir string
	Stdout  bool
}

func (p ShellPatch) Apply() error {
	var cmd *exec.Cmd
	if p.Args != nil {
		cmd = exec.Command(p.Command, p.Args...)
	} else {
		cmd = exec.Command(p.Command)
	}

	if p.WorkDir != "" {
		cmd.Dir = p.WorkDir
	}

	if p.Stdout {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}

func (p ShellPatch) Summary(ctx *context.Context) (string, error) {
	ctx = context.Get(ctx)

	var str strings.Builder

	str.WriteString(ctx.GetPrefix() + ctx.GetCounter() + "shell " + p.Command)
	for _, v := range p.Args {
		str.WriteRune(' ')
		str.WriteString(v)
	}
	str.WriteRune('\n')

	if p.WorkDir != "" {
		str.WriteString(ctx.GetPrefix() + "Workdir: ")
		str.WriteString(p.WorkDir)
		str.WriteRune('\n')
	}

	return str.String(), nil
}
