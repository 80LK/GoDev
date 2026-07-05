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

	validated bool
}

func (p *WriteFilePatch) validate() {
	if p.validated {
		return
	}
	hasCRLF := bytes.Contains(p.OldData, []byte("\r\n"))

	if hasCRLF {
		p.NewData = bytes.ReplaceAll(
			p.NewData,
			[]byte("\n"),
			[]byte("\r\n"),
		)
	}
	p.validated = true
}

func (p WriteFilePatch) Apply() error {
	p.validate()
	return os.WriteFile(
		p.Path,
		p.NewData,
		p.Perm,
	)
}
func (p WriteFilePatch) Diff() (string, error) {
	p.validate()
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
	p.validate()
	switch {
	case len(p.OldData) == 0:
		return "create file " + p.Path

	case bytes.Equal(p.OldData, p.NewData):
		return ""

	default:
		return "modify file " + p.Path
	}
}
