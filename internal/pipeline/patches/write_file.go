package patches

import (
	"bytes"
	"os"

	"github.com/pmezard/go-difflib/difflib"
)

type WriteFilePatch struct {
	Path string

	OldData []byte
	NewData []byte

	Perm os.FileMode
}

func (p WriteFilePatch) Apply() error {
	return os.WriteFile(
		p.Path,
		p.NewData,
		p.Perm,
	)
}
func (p WriteFilePatch) Diff() (string, error) {
	diff := difflib.UnifiedDiff{
		A: difflib.SplitLines(string(p.OldData)),
		B: difflib.SplitLines(string(p.NewData)),

		FromFile: p.Path,
		ToFile:   p.Path,

		Context: 3,
	}

	return difflib.GetUnifiedDiffString(diff)
}
func (p WriteFilePatch) Summary() string {
	switch {
	case len(p.OldData) == 0:
		return "create file " + p.Path

	case bytes.Equal(p.OldData, p.NewData):
		return "unchanged file " + p.Path

	default:
		return "modify file " + p.Path
	}
}
