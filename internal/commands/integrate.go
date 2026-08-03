package commands

import (
	"github.com/80LK/godev/cli"
	"github.com/spf13/cobra"
)

var IntCmd = &cobra.Command{
	Use:     "integrate",
	Aliases: []string{"int"},
	Short:   "Integrate tool in exsist project",
	Long:    "Integrate tool in exsist project",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := cli.IntegrateOptions{}
		opts.DryRun = dryRun

		var err error
		opts.Author, err = cmd.Flags().GetString(_AUTHOR_FLAG)
		if err != nil {
			return err
		}

		opts.Version, err = cmd.Flags().GetString(_VERSION_FLAG)
		if err != nil {
			return err
		}

		return cli.Integrate(opts)
	},
}

func init() {
	flags := IntCmd.Flags()

	flags.StringP(_AUTHOR_FLAG, _AUTHOR_FLAG_P, "", _AUTHOR_FLAG_U)
	flags.StringP(_VERSION_FLAG, _VERSION_FLAG_P, "", _VERSION_FLAG_U)

	Root.AddCommand(IntCmd)
}
