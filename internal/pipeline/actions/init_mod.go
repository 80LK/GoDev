package actions

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline"
	"golang.org/x/mod/modfile"
)

type InitMod struct{}

const TOOL_PATH = "github.com/80LK/godev/cmd/god"
const TOOL_VERSION = "v0.0.0"

func (f InitMod) Plan(ctx *pipeline.Context) ([]pipeline.Patch, error) {
	if ctx.Project == nil {
		return nil, fmt.Errorf("project not found")
	}

	if ctx.Mod != nil {
		return nil, nil
	}

	ctx.Mod = new(modfile.File)
	ctx.Mod.AddModuleStmt(ctx.Project.Project.GetGoModStmt())
	ctx.Mod.AddGoStmt(ctx.GoVer.StringWithoutSuffix())
	ctx.Mod.AddNewRequire(TOOL_PATH, TOOL_VERSION, true)
	ctx.Mod.AddTool(TOOL_PATH)

	return nil, nil
}
