package commands

import (
	"errors"

	"github.com/80LK/godev/cli"
	"github.com/spf13/cobra"

	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
)

var (
	release  bool
	noCommit bool
)

var VersionCmd = &cobra.Command{
	Use:   "version <patch|minor|major>",
	Short: "Bump version in go.project",
	Long: `Bump version in go.project

Args:
  <patch|minor|major> 	Bump patch, minor or major version
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := projectAct.ErrBump
		if len(args) != 1 {
			return err
		}

		opts := cli.VersionBumpOptions{}
		opts.Release = release
		opts.NoCommit = noCommit
		opts.DryRun = dryRun

		opts.Bump, err = projectAct.ToBump(args[0])
		if err != nil {
			return err
		}

		opts.Pre, err = cmd.Flags().GetString("pre")
		if err != nil {
			return err
		}

		return cli.VersionBump(opts)
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
		err := errors.New("Expose pre-release tag")
		if len(args) != 1 {
			return err
		}

		opts := cli.VersionPreReleaseOptions{}
		opts.DryRun = dryRun
		opts.NoCommit = noCommit
		opts.Release = release
		opts.Pre = args[0]

		return cli.VersionPreRelease(opts)
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

		opts := cli.VersionSetOptions{}
		opts.DryRun = dryRun
		opts.NoCommit = noCommit
		opts.Release = release
		opts.Version = args[0]

		return cli.VersionSet(opts)
	},
}

func init() {
	VersionCmd.Flags().StringP("pre", "p", "", "Set pre-release tag")
	VersionCmd.PersistentFlags().BoolVarP(&release, "release", "r", false, "add git tag for release")
	VersionCmd.PersistentFlags().BoolVarP(&noCommit, "no-commit", "c", false, "change version without commit")
	VersionCmd.AddCommand(
		VersionSetCmd,
		VersionPreCmd,
	)
	Root.AddCommand(VersionCmd)
}
