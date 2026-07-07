package patches

import (
	"bytes"
	"os"

	"github.com/pmezard/go-difflib/difflib"
)

func NewWriteFilePatch(path string, oldData, newData []byte, perm os.FileMode) writeFilePatch {
	hasCRLF := bytes.Contains(oldData, []byte("\r\n"))
	if hasCRLF {
		newData = bytes.ReplaceAll(
			newData,
			[]byte("\n"),
			[]byte("\r\n"),
		)
	}
	return writeFilePatch{
		Path:    path,
		OldData: oldData,
		NewData: newData,
		Perm:    perm,
	}
}

type writeFilePatch struct {
	Path string

	OldData []byte
	NewData []byte

	Perm os.FileMode
}

func (p writeFilePatch) Apply() error {
	return os.WriteFile(
		p.Path,
		p.NewData,
		p.Perm,
	)
}
func (p writeFilePatch) Diff() (string, error) {
	diff := difflib.UnifiedDiff{
		A: difflib.SplitLines(string(p.OldData)),
		B: difflib.SplitLines(string(p.NewData)),

		FromFile: p.Path,
		ToFile:   p.Path,

		Context: 3,
	}

	return difflib.GetUnifiedDiffString(diff)
}
func (p writeFilePatch) Summary() string {
	switch {
	case len(p.OldData) == 0:
		return "create file " + p.Path

	case bytes.Equal(p.OldData, p.NewData):
		return ""

	default:
		return "modify file " + p.Path
	}
}
