package errors

import (
	"fmt"
)

type notFile struct {
	path string
}

func (n notFile) Error() string {
	return fmt.Sprintf("%s not file", n.path)
}

type notDir struct {
	path string
}

func (n notDir) Error() string {
	return fmt.Sprintf("%s not directory", n.path)
}

func ErrNotFile(path string) error {
	return notFile{path: path}
}
func ErrNotDir(path string) error {
	return notFile{path: path}
}
