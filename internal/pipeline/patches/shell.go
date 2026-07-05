package patches

import (
	"os"
	"os/exec"
	"strings"
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

func (p ShellPatch) Diff() (string, error) {
	return "", nil
}

func (p ShellPatch) Summary() string {
	var str strings.Builder

	str.WriteString("shell ")
	str.WriteString(p.Command)
	for _, v := range p.Args {
		str.WriteRune(' ')
		str.WriteString(v)
	}

	if p.WorkDir != "" {
		str.WriteRune('\n')
		str.WriteString("Workdir: ")
		str.WriteString(p.WorkDir)
	}

	return str.String()
}
