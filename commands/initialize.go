package commands

import (
	"path"
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/project"

	fsAct "github.com/80LK/godev/internal/pipeline/actions/fs"
	gitAct "github.com/80LK/godev/internal/pipeline/actions/git"
	goAct "github.com/80LK/godev/internal/pipeline/actions/go"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
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
		ctx := context.New(dryRun)

		moduleName := args[0]
		template, err := cmd.Flags().GetString(_TEMPLATE_FLAG)
		if err != nil {
			return err
		}

		force, err := cmd.Flags().GetBool(_FORCE_FLAG)
		if err != nil {
			return err
		}
		author, err := cmd.Flags().GetString(_AUTHOR_FLAG)
		if err != nil {
			return err
		}

		version, err := cmd.Flags().GetString(_VERSION_FLAG)
		if err != nil {
			return err
		}

		//ctx.ProjectDir now CWD
		if moduleName == "." {
			moduleName = filepath.Base(ctx.ProjectDir)
		} else {
			ctx.ProjectDir = filepath.Join(ctx.ProjectDir, path.Base(moduleName))
		}

		pl := pipeline.New()
		if force {
			pl.Add(fsAct.EnsureDir{Path: ctx.ProjectDir, Perm: 0777})
		} else {
			pl.Add(fsAct.EnsureEmptyDir{Path: ctx.ProjectDir, Perm: 0777})
		}

		pl.Add(
			fsAct.CheckNotExists{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			fsAct.CheckNotExists{Path: project.GetGoModFile(ctx.ProjectDir)},

			projectAct.InitProjectContext{},

			projectAct.InitProject{
				ModuleName: moduleName,
				Author:     author,
			},
			projectAct.InitMod{},
		)

		if version != "" {
			pl.Add(projectAct.VersionSet{Value: version})
		}

		return pl.Add(
			projectAct.CreateFromTemplate{
				Template: template,
			},

			gitAct.GitInit{},
			goAct.Tidy{},
			gitAct.GitCommit{Value: "init commit"},
		).Execute(ctx)
	},
}

func init() {
	flags := InitCmd.Flags()

	flags.StringP(_TEMPLATE_FLAG, "t", "app", "usage template. Default: app. Available: app; module.")
	flags.BoolP(_FORCE_FLAG, "f", false, "force initialize project in non-empty directory")
	flags.StringP(_AUTHOR_FLAG, _AUTHOR_FLAG_P, "", _AUTHOR_FLAG_U)
	flags.StringP(_VERSION_FLAG, _VERSION_FLAG_P, "", _VERSION_FLAG_U)

	Root.AddCommand(InitCmd)
}
