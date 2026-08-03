package commands

import (
	"errors"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/project"
	"github.com/spf13/cobra"

	fsAct "github.com/80LK/godev/internal/pipeline/actions/fs"
	gitAct "github.com/80LK/godev/internal/pipeline/actions/git"
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
		if len(args) != 1 {
			return projectAct.ErrBump
		}

		bump, err := projectAct.ToBump(args[0])
		if err != nil {
			return err
		}

		pre, err := cmd.Flags().GetString("pre")
		if err != nil {
			return err
		}

		ctx := context.New(dryRun)
		pl := pipeline.New().Add(
			fsAct.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
		)
		if !noCommit {
			pl.Add(gitAct.CheckClearGit{})
		}
		pl.Add(
			projectAct.InitProjectContext{},
			projectAct.RunScript{IgnoreNotFound: true, Name: project.LifecycleName(project.PhaseBefore, "version")},
			projectAct.VersionBump{OldVersionKey: "old_version", NewVersionKey: "new_version", Value: bump},
		)
		if pre != "" {
			pl.Add(projectAct.VersionSetPre{NewVersionKey: "new_version", Value: pre})
		}

		pl.Add(
			projectAct.EncodeGoProject{},
			projectAct.PatchSources{OldVersionKey: "old_version"},
			projectAct.GenerateMeta{},
		)

		if !noCommit {
			pl.Add(gitAct.GitCommit{InputKey: "new_version"})
		}

		if !noCommit && release {
			pl.Add(gitAct.GitTagVersion{})
		}
		pl.Add(projectAct.RunScript{IgnoreNotFound: true, Name: project.LifecycleName(project.PhaseAfter, "version")})

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

		pl := pipeline.New().Add(
			fsAct.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
		)

		if !noCommit {
			pl.Add(gitAct.CheckClearGit{})
		}

		pl.Add(
			projectAct.InitProjectContext{},
			projectAct.VersionSetPre{OldVersionKey: "old_version", NewVersionKey: "new_version", Value: args[0]},
			projectAct.EncodeGoProject{},
			projectAct.GenerateMeta{},
		)
		if !noCommit {
			pl.Add(gitAct.GitCommit{InputKey: "new_version"})
		}

		if !noCommit && release {
			pl.Add(gitAct.GitTagVersion{})
		}

		return pl.Execute(ctx)
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

		pl := pipeline.New().Add(
			fsAct.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
		)
		if !noCommit {
			pl.Add(gitAct.CheckClearGit{})
		}
		pl.Add(
			projectAct.InitProjectContext{},
			projectAct.VersionSet{OldVersionKey: "old_version", NewVersionKey: "new_version", Value: args[0]},
			projectAct.EncodeGoProject{},
			projectAct.PatchSources{OldVersionKey: "old_version"},
			projectAct.GenerateMeta{},
		)

		if !noCommit {
			pl.Add(gitAct.GitCommit{InputKey: "new_version"})
		}

		if !noCommit && release {
			pl.Add(gitAct.GitTagVersion{})
		}
		return pl.Execute(ctx)
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
