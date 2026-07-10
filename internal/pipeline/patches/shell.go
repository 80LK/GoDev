package patches

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"

	"github.com/80LK/godev/internal/pipeline/patches/context"
)

type ShellPatch struct {
	Command string
	Args    []string
	WorkDir string
	Env     []string

	Stdout bool
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

	if p.Env != nil {
		cmd.Env = slices.Concat(
			os.Environ(),
			p.Env,
		)
	} else {
		cmd.Env = os.Environ()
	}

	if p.Stdout {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error shell %q: %s", p.Command+" "+strings.Join(p.Args, " "), err)
	}

	return nil
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
