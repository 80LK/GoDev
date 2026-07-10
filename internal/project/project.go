package project

import (
	"github.com/80LK/godev/internal/version"
)

type GoProject struct {
	Project *ProjectInfo          `modlike:"project"`
	Meta    bool                  `modlike:"meta"`
	Builds  map[string]*BuildInfo `modlike:"build"`
	Scripts map[string]*Script    `modlike:"script"`
}

func New() *GoProject {
	return &GoProject{
		Project: &ProjectInfo{
			Version: &version.Version{},
		},
	}
}
