package commands

import (
	"os"
	"path"
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/actions"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "initialize <module-name> [template]",
	Aliases: []string{"init"},
	Short:   "Initialize a new module",
	Long: `Initialize a new module
	
Args:
  <module-name >	Name of the module (required)
                	Use "." to initialize in the current directory
                	In this case, the module name is derived from the directory name
  [  template  ]	Usage template. Default: app. Available: app; module.`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := pipeline.NewContext(dryRun)

		moduleName := args[0]
		var template string
		if len(args) == 2 {
			template = args[1]
		} else {
			template = "app"
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

			actions.LoadProject{},

			actions.InitProject{
				ModuleName: moduleName,
			},

			actions.CreateFromTemplate{
				Template: template,
			},
		)

		return pl.Execute(ctx)
	},
}

func init() {
	Root.AddCommand(initCmd)
}
