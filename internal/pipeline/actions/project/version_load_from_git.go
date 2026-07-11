package actions

import (
	"math"
	"os/exec"
	"strconv"
	"strings"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"
	"github.com/80LK/godev/internal/version"
)

type VersionLoadFromGit struct{}

func (v VersionLoadFromGit) Plan(ctx *context.Context) ([]patches.Patch, error) {
	cmd := exec.Command("git", "for-each-ref", "--merged", "HEAD", "--format=%(refname:short)", "refs/tags")
	cmd.Dir = ctx.ProjectDir

	data, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	tags := strings.Fields(string(data))
	var minRange uint64 = math.MaxUint64
	var closestVersion *version.Version
	for _, tag := range tags {
		ver, err := version.Parse(tag)
		if err != nil {
			continue
		}

		if ver.Major != ctx.GoProject.Project.Version.Major {
			continue
		}

		cmd := exec.Command("git", "rev-list", "--count", ver.String()+"..HEAD")
		cmd.Dir = ctx.ProjectDir

		data, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		dist, err := strconv.ParseUint(strings.TrimSpace(string(data)), 10, 64)
		if err != nil {
			return nil, err
		}

		if dist < minRange {
			closestVersion = ver
			minRange = dist
		}
	}

	if closestVersion != nil {
		ctx.GoProject.Project.Version = closestVersion
	}

	return nil, nil
}
