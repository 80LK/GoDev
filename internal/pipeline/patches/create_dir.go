package patches

import (
	"os"
)

type CreateDirPatch struct {
	Path string
	Perm os.FileMode
}

func (p CreateDirPatch) Apply() error {
	return os.MkdirAll(p.Path, p.Perm)
}

func (p CreateDirPatch) Diff() (string, error) {
	return "", nil
}

func (p CreateDirPatch) Summary() string {
	return "mkdir " + p.Path
}
