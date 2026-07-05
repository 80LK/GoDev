package actions

import (
	"errors"
	"strings"

	"github.com/80LK/godev/internal/pipeline"
)

type Bump string

var ErrBump = errors.New("expose patch, minor or major")

const (
	BUMP_PATCH Bump = "patch"
	BUMP_MINOR Bump = "minor"
	BUMP_MAJOR Bump = "major"
)

func ToBump(str string) (Bump, error) {
	str = strings.ToLower(str)

	switch str {
	case string(BUMP_PATCH):
		return BUMP_PATCH, nil

	case string(BUMP_MINOR):
		return BUMP_MINOR, nil

	case string(BUMP_MAJOR):
		return BUMP_MAJOR, nil
	default:
		return "", ErrBump
	}
}

type VersionBump struct {
	Value Bump
}

func (v VersionBump) Plan(ctx *pipeline.Context) ([]pipeline.Patch, error) {
	if v.Value == "" {
		return nil, ErrBump
	}

	switch v.Value {
	case BUMP_PATCH:
		ctx.GoProject.Project.Version.Patch++
	case BUMP_MINOR:
		ctx.GoProject.Project.Version.Minor++
	case BUMP_MAJOR:
		ctx.GoProject.Project.Version.Major++
	}

	return nil, nil
}
