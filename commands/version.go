package commands

import (
	"errors"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/actions"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/internal/project"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version <patch|minor|major>",
	Short: "Bump version in go.project",
	Long: `Bump version in go.project

Args:
  <patch|minor|major> 	Bump patch, minor or major version
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return actions.ErrBump
		}

		bump, err := actions.ToBump(args[0])
		if err != nil {
			return err
		}

		pre, err := cmd.Flags().GetString("pre")
		if err != nil {
			return err
		}

		ctx := context.New(dryRun)
		pl := pipeline.New().Add(
			actions.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			actions.InitProjectContext{},
			actions.VersionBump{Value: bump},
		)
		if pre != "" {
			pl.Add(actions.VersionSetPre{Value: pre})
		}

		pl.Add(
			actions.EncodeGoProject{},
		)

		return pl.Execute(ctx)
	},
}

var VersionPreCmd = &cobra.Command{
	Use:   "pre <pre-release-tag>",
	Short: "Set pre-release tag in go.project",
	Long: `Set pre-release tag in go.project.

Args:
  <pre-release-tag> 	Setted pre-release tag in go.project
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Expose pre-release tag")
		}

		ctx := context.New(dryRun)

		return pipeline.New().Add(
			actions.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			actions.InitProjectContext{},
			actions.VersionSetPre{Value: args[0]},
			actions.EncodeGoProject{},
		).Execute(ctx)
	},
}

var VersionSetCmd = &cobra.Command{
	Use:   "set <version>",
	Short: "Set version in go.project",
	Long: `Set version tag in go.project.

Args:
  <version> 	Setted version in go.project
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Expose version")
		}

		ctx := context.New(dryRun)

		return pipeline.New().Add(
			actions.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			actions.InitProjectContext{},
			actions.VersionSet{Value: args[0]},
			actions.EncodeGoProject{},
		).Execute(ctx)
	},
}

func init() {
	VersionCmd.Flags().StringP("pre", "p", "", "Set pre-release tag")
	VersionCmd.AddCommand(
		VersionSetCmd,
		VersionPreCmd,
	)
	Root.AddCommand(VersionCmd)
}
