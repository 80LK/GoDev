package project

import (
	"strconv"

	"github.com/80LK/godev/internal/version"
)

type ProjectInfo struct {
	Name    string           `modlike:"name" json:"name"`
	Version *version.Version `modlike:"version" json:"version"`
	Module  string           `modlike:"module" json:"module"`
	Author  string           `modlike:"author" json:"author,omitempty"`
}

func (p *ProjectInfo) GetGoModStmt() string {
	if p.Version.Major < 2 {
		return p.Module
	} else {
		return p.Module + "/v" + strconv.FormatUint(uint64(p.Version.Major), 10)
	}
}
