package project

import (
	path "path/filepath"
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

type GoProject struct {
	Project *ProjectInfo `modlike:"project"`
}

func Parse(doc modlike.Document) (*GoProject, error) {
	var goProj GoProject
	err := doc.Decode(&goProj)
	if err != nil {
		return nil, err
	}
	return &goProj, nil
}

func ParseFromGoMod(file *modfile.File) (*GoProject, error) {
	moduleParts := strings.Split(file.Module.Mod.Path, "/")
	modulePartsLen := len(moduleParts)

	var goProject GoProject = GoProject{
		Project: &ProjectInfo{},
	}
	lastPart := moduleParts[modulePartsLen-1]

	if modulePartsLen > 1 && version.IsValid(moduleParts[modulePartsLen-1]) {
		goProject.Project.Version, _ = version.Parse(lastPart)
		moduleParts = moduleParts[:modulePartsLen-1]
		modulePartsLen--
		lastPart = moduleParts[modulePartsLen-1]
	}

	if modulePartsLen == 1 {
		goProject.Project.Name = moduleParts[0]
	} else {
		goProject.Project.Author = moduleParts[0]
		goProject.Project.Name = lastPart
	}

	goProject.Project.Module = strings.Join(moduleParts, "/")

	return &goProject, nil
}

func GetGoProjectFile(dir string) string {
	return path.Join(dir, "go.project")
}

func GetGoModFile(dir string) string {
	return path.Join(dir, "go.mod")
}
