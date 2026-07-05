package project

import (
	path "path/filepath"
	"strconv"
	"strings"

	"github.com/80LK/godev/internal/version"

	"github.com/80lk/modlike"
	"golang.org/x/mod/modfile"
)

type ProjectInfo struct {
	Name    string           `modlike:"name"`
	Version *version.Version `modlike:"version"`
	Module  string           `modlike:"module"`
	Author  string           `modlike:"author"`
}

func (p *ProjectInfo) GetGoModStmt() string {
	if p.Version.Major < 2 {
		return p.Module
	} else {
		return p.Module + "/v" + strconv.FormatUint(uint64(p.Version.Major), 10)
	}
}

type GoProject struct {
	Project *ProjectInfo `modlike:"project"`
}

func New() *GoProject {
	return &GoProject{
		Project: &ProjectInfo{
			Version: &version.Version{},
		},
	}
}

func ParseIn(doc modlike.Document, out *GoProject) error {
	err := doc.Decode(out)
	if err != nil {
		return err
	}
	return nil
}

func ParseFromGoModIn(file *modfile.File, out *GoProject) error {
	return ParseFromGoModuleNameIn(file.Module.Mod.Path, out)
}

func GetGoProjectFile(dir string) string {
	return path.Join(dir, "go.project")
}

func GetGoModFile(dir string) string {
	return path.Join(dir, "go.mod")
}

func ParseFromGoModuleNameIn(name string, out *GoProject) error {
	moduleParts := strings.Split(name, "/")
	modulePartsLen := len(moduleParts)

	lastPart := moduleParts[modulePartsLen-1]

	if modulePartsLen > 1 && version.IsValid(moduleParts[modulePartsLen-1]) {
		out.Project.Version, _ = version.Parse(lastPart)
		moduleParts = moduleParts[:modulePartsLen-1]
		modulePartsLen--
		lastPart = moduleParts[modulePartsLen-1]
	}

	switch modulePartsLen {
	case 1:
		out.Project.Name = moduleParts[0]
	case 2:
		out.Project.Author = moduleParts[0]
		out.Project.Name = lastPart
	default:
		out.Project.Author = moduleParts[1]
		out.Project.Name = lastPart
	}

	out.Project.Module = strings.Join(moduleParts, "/")

	return nil
}
