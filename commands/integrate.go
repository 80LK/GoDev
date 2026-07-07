package commands

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/actions"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/project"
	"github.com/spf13/cobra"
)

var IntCmd = &cobra.Command{
	Use:     "integrate",
	Aliases: []string{"int"},
	Short:   "Integrate tool in exsist project",
	Long:    "Integrate tool in exsist project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.New(dryRun)

		author, err := cmd.Flags().GetString(_AUTHOR_FLAG)
		if err != nil {
			return err
		}

		version, err := cmd.Flags().GetString(_VERSION_FLAG)
		if err != nil {
			return err
		}

		pl := pipeline.New().Add(
			actions.CheckNotExists{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			actions.CheckExistsFile{Path: project.GetGoModFile(ctx.ProjectDir)},

			actions.InitProjectContext{},

			actions.IntMod{},
			actions.IntProject{
				Author: author,
			},
		)

		if version != "" {
			pl.Add(actions.VersionSet{Value: version})
		}

		return pl.Add(
			actions.EncodeGoMod{},
			actions.EncodeGoProject{},

			actions.GitInit{},
		).Execute(ctx)
	},
}

func init() {
	flags := IntCmd.Flags()

	flags.StringP(_AUTHOR_FLAG, _AUTHOR_FLAG_P, "", _AUTHOR_FLAG_U)
	flags.StringP(_VERSION_FLAG, _VERSION_FLAG_P, "", _VERSION_FLAG_U)

	Root.AddCommand(IntCmd)
}
