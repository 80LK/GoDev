package commands

import (
	"os"
	"path"
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/actions"

	"github.com/spf13/cobra"
)

const _TEMPLATE_FLAG = "template"

var initCmd = &cobra.Command{
	Use:     "initialize <module-name>",
	Aliases: []string{"init"},
	Short:   "Initialize a new module",
	Long: `Initialize a new module

Args:
  <module-name >	Name of the module (required)
                	Use "." to initialize in the current directory
                	In this case, the module name is derived from the directory name`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := pipeline.NewContext(dryRun)

		moduleName := args[0]
		template, err := cmd.Flags().GetString(_TEMPLATE_FLAG)
		if err != nil {
			return err
		}

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

		pl := pipeline.New().Add(
			actions.EnsureEmptyDir{Path: ctx.ProjectDir, Perm: 0777},

			actions.InitProjectContext{},

			actions.InitProject{
				ModuleName: moduleName,
			},
			actions.InitMod{},

			actions.CreateFromTemplate{
				Template: template,
			},
		)

		return pl.Execute(ctx)
	},
}

func init() {
	initCmd.Flags().StringP(_TEMPLATE_FLAG, "t", "app", "usage template. Default: app. Available: app; module.")
	Root.AddCommand(initCmd)
}
