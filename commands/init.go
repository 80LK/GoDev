package commands

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/actions"
	"github.com/80LK/godev/internal/project"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <module-name>",
	Short: "init module",
	Long: `Initialize a new module or integrate into an existing module
	
Args:
  <module-name> 	Name of the module (required)
                	Use "." to initialize in the current directory
                	In this case, the module name is derived from the directory name`,
	Args: func(cmd *cobra.Command, args []string) error {
		largs := len(args)
		if largs < 1 {
			return fmt.Errorf("Module name is required")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := pipeline.NewContext(dryRun)
		moduleName := args[0]

		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		if moduleName == "." {
			ctx.ProjectDir = cwd
			moduleName = filepath.Base(ctx.ProjectDir)
		} else {
			ctx.ProjectDir = filepath.Join(cwd, path.Base(moduleName))
		}

		const (
			GO_PROJECT_KEY = "go.project"
			GO_MOD_KEY     = "go.mod"
		)

		pl := pipeline.Pipeline{
			Actions: []pipeline.Action{
				actions.EnsureDir{
					Path: ctx.ProjectDir,
					Perm: 0777,
				},

				actions.CheckNotExists{
					Path: project.GetGoProjectFile(ctx.ProjectDir),
				},

				actions.LoadProject{},

				actions.InitProject{
					ModuleName: moduleName,
				},

				actions.EncodeGoProject{
					OutputKey: GO_PROJECT_KEY,
				},
				actions.WriteFile{
					InputKey: GO_PROJECT_KEY,
					Path:     project.GetGoProjectFile(ctx.ProjectDir),
					Perm:     0666,
				},
				actions.ConditionAction{
					Condition: func(ctx *pipeline.Context) bool {
						return !ctx.HasGoMod
					},
					Action: pipeline.Pipeline{
						Actions: []pipeline.Action{
							actions.EncodeGoMod{
								OutputKey: GO_MOD_KEY,
							},
							actions.WriteFile{
								InputKey: GO_MOD_KEY,
								Path:     project.GetGoModFile(ctx.ProjectDir),
								Perm:     0666,
							},
						},
					},
				},
				actions.EnsureDir{
					Path: filepath.Join(ctx.ProjectDir, "internal"),
					Perm: 0777,
				},
				actions.ConditionAction{
					Condition: func(ctx *pipeline.Context) bool {
						return !ctx.HasGoProject && !ctx.HasGoMod
					},
					Action: actions.WriteFile{
						Value: fmt.Appendf(nil, "package %s", moduleName),
						Path:  filepath.Join(ctx.ProjectDir, path.Base(moduleName)+".go"),
						Perm:  0666,
					},
				},
			},
		}

		return pipeline.Execute(ctx, &pl)
	},
}

func init() {

	Root.AddCommand(initCmd)
}
