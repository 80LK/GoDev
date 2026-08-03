package project

import (
	"github.com/80LK/godev/internal/version"
)

type GoProject struct {
	Project       *ProjectInfo           `modlike:"project" json:"project"`
	Meta          bool                   `modlike:"meta,omitempty" json:"meta,omitempty"`
	Builds        map[string]*BuildInfo  `modlike:"build" json:"build,omitempty"`
	Scripts       map[string]*RootScript `modlike:"script" json:"script,omitempty"`
	ConfigVersion uint                   `json:"config_version"`
}

func New() *GoProject {
	return &GoProject{
		Project: &ProjectInfo{
			Version: &version.Version{},
		},
	}
}
