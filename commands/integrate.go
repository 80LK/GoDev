package commands

import (
	"os"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/actions"
	"github.com/80LK/godev/internal/project"
	"github.com/spf13/cobra"
)

var intgCmd = &cobra.Command{
	Use:     "integrate",
	Aliases: []string{"int"},
	Short:   "Integrate tool in exsist project",
	Long:    "Integrate tool in exsist project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := pipeline.NewContext(dryRun)
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		ctx.ProjectDir = cwd

		return pipeline.New().Add(
			actions.CheckNotExists{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			actions.CheckExistsFile{Path: project.GetGoModFile(ctx.ProjectDir)},

			actions.InitProjectContext{},

			actions.IntMod{},
			actions.IntProject{},

			actions.EncodeGoMod{},
			actions.EncodeGoProject{},

			actions.GitInit{},
		).Execute(ctx)
	},
}

func init() {
	Root.AddCommand(intgCmd)
}
