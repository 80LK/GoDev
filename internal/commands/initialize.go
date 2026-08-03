package commands

import (
	"github.com/80LK/godev/cli"

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
		opts := cli.InitializeOptions{}
		opts.DryRun = dryRun

		opts.ModuleName = args[0]
		var err error
		opts.Template, err = cmd.Flags().GetString(_TEMPLATE_FLAG)
		if err != nil {
			return err
		}

		opts.Force, err = cmd.Flags().GetBool(_FORCE_FLAG)
		if err != nil {
			return err
		}
		opts.Author, err = cmd.Flags().GetString(_AUTHOR_FLAG)
		if err != nil {
			return err
		}

		opts.Version, err = cmd.Flags().GetString(_VERSION_FLAG)
		if err != nil {
			return err
		}

		return cli.Initialize(opts)
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
