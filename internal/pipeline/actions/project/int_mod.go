package actions

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/pipeline/patches"

	"github.com/80LK/godev/internal/utils"
)

type IntMod struct{}

func (i IntMod) Plan(ctx *context.Context) ([]patches.Patch, error) {
	if ctx.Mod == nil {
		return nil, fmt.Errorf("go.mod not found")
	}

	if !utils.HasTool(ctx.Mod, TOOL_PATH) {
		ctx.Mod.AddTool(TOOL_PATH)
	}

	return nil, nil
}
