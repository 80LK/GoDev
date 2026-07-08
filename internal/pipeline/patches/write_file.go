package patches

import (
	"bytes"
	"os"
	"strings"

	"github.com/80LK/godev/internal/pipeline/patches/context"
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

func (p writeFilePatch) Summary(ctx *context.Context) (string, error) {
	if bytes.Equal(p.OldData, p.NewData) {
		return "", nil
	}

	ctx = context.Get(ctx)

	diff := difflib.UnifiedDiff{
		A: difflib.SplitLines(string(p.OldData)),
		B: difflib.SplitLines(string(p.NewData)),

		FromFile: p.Path,
		ToFile:   p.Path,

		Context: 3,
	}

	diffS, err := difflib.GetUnifiedDiffString(diff)
	if err != nil {
		return "", err
	}

	lines := difflib.SplitLines(diffS)
	var str strings.Builder
	str.WriteString(ctx.GetPrefix() + ctx.GetCounter())
	if len(p.OldData) == 0 {
		str.WriteString("create file ")
	} else {
		str.WriteString("modify file ")
	}
	str.WriteString(p.Path + "\n")

	for _, line := range lines {
		str.WriteString(ctx.GetPrefix() + line)
	}

	return str.String(), nil
}
