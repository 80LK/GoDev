package actions

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/utils"
)

type IntMod struct{}

func (i IntMod) Plan(ctx *pipeline.Context) ([]pipeline.Patch, error) {
	if ctx.Mod == nil {
		return nil, fmt.Errorf("go.mod not found")
	}

	if !utils.HasTool(ctx.Mod, TOOL_PATH) {
		ctx.Mod.AddTool(TOOL_PATH)
	}

	if !utils.HasRequire(ctx.Mod, TOOL_PATH) {
		ctx.Mod.AddNewRequire(TOOL_PATH, TOOL_VERSION, true)
	}
	return nil, nil
}
